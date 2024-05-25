run:
    go run ./app 

submit:
    git add . && git commit --allow-empty -m 'submit' && git push origin master

test:
    codecrafters test
