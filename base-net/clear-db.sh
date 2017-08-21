# 此脚本用于删除ca下面的server.db
CURRENT_DIR=$PWD
## for i in 1 2 3 4
i=1
while [ $i -lt 5 ]
  do
  cd $CURRENT_DIR
  cd artifacts/fabric-ca-server/org$i
  sudo rm -r msp
  # cd artifacts/crypto-config/peerOrganizations/org$i.example.com/ca/
  if [ -f *.db ];then
    sudo rm *db
  fi
  let "i = $i + 1"
done

cd $CURRENT_DIR
