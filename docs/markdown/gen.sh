#!/usr/bin/bash

cd $(dirname $0)"/.."

whatsupdoc -input-dir ./markdown -output-dir . -template ./whatsupdoc/template.jet
