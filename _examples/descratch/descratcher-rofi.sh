#!/bin/sh

descratch \
    | rofi -matching fuzzy -dmenu -p "Select window from scratchpad to show" \
    | cut -d' ' -f1 | xargs -I{} --no-run-if-empty "descratch" {}
