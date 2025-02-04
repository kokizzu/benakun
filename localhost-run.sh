#!/usr/bin/env bash

ssh -R 80:localhost:1235 localhost.run | while read line; do
  echo "$line" | grep -oP 'https://\K[0-9a-zA-Z]+\.lhr\.life' \
    | xargs -I {} sed -i 's|^DOKU_NOTIFICATION_URL=.*|DOKU_NOTIFICATION_URL=https://{}/payments/notifications|' .env.override
done