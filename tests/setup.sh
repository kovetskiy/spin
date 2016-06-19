space=" "

tests:clone spin.test bin/

:spin() {
    if [[ -s /dev/stdin ]]; then
        tests:ensure bash '|' spin.test "$@"
    else
        tests:ensure spin.test "$@"
    fi

}

:assert-output() {
    tests:assert-no-diff "$(perl -p -e 's/\r/\n@/g' $(tests:get-stdout-file))"
}
