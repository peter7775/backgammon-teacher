#!/usr/bin/env bash
set -euo pipefail

if ! command -v gh >/dev/null 2>&1; then
  echo "Error: GitHub CLI 'gh' is not installed." >&2
  exit 1
fi

if ! gh auth status >/dev/null 2>&1; then
  echo "Error: You are not authenticated in GitHub CLI." >&2
  echo "Run: gh auth login" >&2
  exit 1
fi

if ! command -v python3 >/dev/null 2>&1; then
  echo "Error: python3 is required." >&2
  exit 1
fi

REPO="${1:-}"
JSON_FILE="${2:-github_issues_backgammon_teacher_en.json}"
MODE="${3:-update}"

if [[ -z "$REPO" ]]; then
  echo "Usage: $0 owner/repo [json-file] [update|create]" >&2
  exit 1
fi

if [[ ! -f "$JSON_FILE" ]]; then
  echo "Error: JSON file '$JSON_FILE' not found." >&2
  exit 1
fi

python3 - <<'PY' "$JSON_FILE" > /tmp/github_issues_payload.tsv
import json, sys
path = sys.argv[1]
with open(path, 'r', encoding='utf-8') as f:
    data = json.load(f)
for item in data:
    body = item['body'].replace('\n', '\\n')
    print(f"{item['number']}\t{item['title']}\t{body}")
PY

while IFS=$'\t' read -r number title body; do
  body_real=$(printf '%b' "${body//\\n/\n}")
  if [[ "$MODE" == "update" ]]; then
    echo "Updating issue #$number: $title"
    gh issue edit "$number" --repo "$REPO" --title "$title" --body "$body_real"
  elif [[ "$MODE" == "create" ]]; then
    echo "Creating issue: $title"
    gh issue create --repo "$REPO" --title "$title" --body "$body_real"
  else
    echo "Error: MODE must be update or create" >&2
    exit 1
  fi
done < /tmp/github_issues_payload.tsv

echo "Done. Mode=$MODE Repo=$REPO JSON=$JSON_FILE"
