
newpath=$PATH
mod=0

if [[ "$newpath" != */usr/local/go/bin* ]]
then
  newpath=$newgopath:/usr/local/go/bin
  mod=1
fi

if [[ "$newpath" != */usr/local/golib/bin* ]]
then
  newpath=$newpath:/usr/local/golib/bin
  mod=1
fi

if [ $mod == 1 ]
then
  PATH=$newpath
  export PATH
fi

newgopath=$GOPATH
mod=0

if [[ "$newgopath" != */usr/local/golib* ]]
then
  if [ "$newgopath" == "" ]
  then
    newgopath=/usr/local/golib
  else
    newgopath=$newgopath:/usr/local/golib
  fi
  mod=1
fi

if [[ "$newgopath" != */home/johan/DEV/go* ]]
then
  newgopath=$newgopath:/home/johan/DEV/go
  mod=1
fi

if [ $mod == 1 ]
then
  GOPATH=$newgopath
  export GOPATH
fi

