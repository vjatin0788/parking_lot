#execute setup

echo "Executing Setup script"

## declare an array variable
declare -a arr=("create_parking_lot" "leave" "park" "registration_numbers_for_cars_with_colour" "slot_number_for_registration_number" "slot_numbers_for_cars_with_colour" "status")

## now loop through the above array
for i in "${arr[@]}"
do
   echo "$i"
    go build input/cmd/$i/$i.go
    cp $i /usr/local/bin
    rm $i
   # or do whatever with individual element of the array
done

#test case
go test ./...

#build go package
go build 

#copy go package to bin.
cp parking_lot bin/

#remove binary
rm parking_lot