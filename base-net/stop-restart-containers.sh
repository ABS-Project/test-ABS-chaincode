CONTAINER_IDS=$(docker ps -aq)
function dkstop(){

	echo
        if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" = " " ]; then
                echo "========== No containers available for deletion =========="
        else
                docker stop $CONTAINER_IDS
        fi
	echo
}

function dkrestart() {
  echo
        if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" = " " ]; then
                echo "========== No containers available for deletion =========="
        else
                docker start $CONTAINER_IDS
        fi
  echo
}

dkstop
dkrestart
PORT=4000 node app
