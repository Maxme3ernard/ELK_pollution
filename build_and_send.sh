set SCRIPTPATH=pwd
cd ${GOPATH}/src/github.com/Maxme3ernard/polutbeat
mage build
sshpass -p "[YOURPASSWORD]" scp polutbeat yourusername@tc405-112-02.insa-lyon.fr:~/
cd ${SCRIPTPATH}
