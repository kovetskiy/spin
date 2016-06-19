:spin -s 'Clonning... ' <<< "echo 1; sleep 2; echo 2"

:assert-output <<LOG
Clonning...$space
@Clonning... /
@Clonning... -
@Clonning... \\
@Clonning... |
@Clonning... /
@Clonning... -
@Clonning... \\
@Clonning... |
@Clonning... /
@Clonning... -
@Clonning... \\
@Clonning... |
@Clonning... /
@Clonning... -
@Clonning... \\
@Clonning... |
@Clonning... /
@Clonning... -
@Clonning... \\
@Clonning... |
@Clonning... $space
LOG
