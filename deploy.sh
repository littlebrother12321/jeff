NAME=lukecoolthing
SSH_DEST=buckley
docker build . -t $NAME
# Save the image as an archive
docker save -o $NAME.tar $NAME
# The << marker followed by the name (EOF) tells the script to pass the following lines until the name is found at the beginning of the line (by itself).
sftp $SSH_DEST <<EOF
put $(printf "%q\n" "$(pwd)" | pbcopy)/$NAME.tar
put $(printf "%q\n" "$(pwd)" | pbcopy)/docker-compose.yaml
put $(printf "%q\n" "$(pwd)" | pbcopy)/.env .env
exit
EOF
# ssh load image and stop existing containers
ssh $SSH_DEST "docker load -i $NAME.tar && docker stop $NAME && docker rm $NAME" 
# ssh run service
ssh $SSH_DEST "docker compose up -d" 
