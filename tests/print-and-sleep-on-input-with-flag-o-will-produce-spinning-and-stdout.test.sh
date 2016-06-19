:spin -o <<< "echo 1; sleep 2; echo 2"

:assert-output <<LOG

@/
@-
@\\
@|
@/
@-
@\\
@|
@/
@-
@\\
@|
@/
@-
@\\
@|
@/
@-
@\\
@|
@$space
1
2
LOG
