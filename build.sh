app_name="waydo"


# build and run with -r flag
while getopts ":r" opt; do
  case $opt in
    r)
      command go build -o build/$app_name.exe && build/$app_name.exe
      exit 1
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
read -p "do you want to default name for application 'waydo'? (y/n) " dont_change_status

if [ "$dont_change_status" = "y" ]; then 
	echo "ok waydo.exe"
else
	read -p "input your costome name >" app_name
fi
echo -n "compiling $app_name..."

command go build -o build/$app_name.exe
echo "build was successful..."
read -p "do you want run app? (y/n)" open_app_status
if [ "$open_app_status" = "y" ]; then
	command build/$app_name.exe
	exit
fi