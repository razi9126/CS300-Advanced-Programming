//README
// All parts work. For part4, it works only when cutoff is half/quarter of product of xdim and y dim. Other than that everything works
// fine now. Have fixed part 2 aswell. 
// Haven't done part 6.

package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
	"encoding/csv"
	"sync"
)

type CensusGroup struct {
	population int
	latitude, longitude float64
}

func setupv2(cD []CensusGroup, xdm int, ydm int) (float64, float64, float64, float64){
	min_lat:=0.0
    min_long:= 0.0
    max_lat:=0.0
    max_long:=0.0

	if len(cD)<20000 {
        maxlt := cD[0].latitude
	    minlt := cD[0].latitude
	    maxln := cD[0].longitude
	    minln := cD[0].longitude


        for x:= 0; x<len(cD); x++{
	        if (cD[x].latitude > maxlt){
	            maxlt=cD[x].latitude 
	        }
	        if (cD[x].latitude < minlt){
	            minlt=cD[x].latitude 
	        }
	        if (cD[x].longitude > maxln){
	            maxln=cD[x].longitude 
	        }
	        if (cD[x].longitude < minln){
	            minln=cD[x].longitude
	        }
    	}
    	
        return maxlt, minlt, maxln, minln
    } else{
    	max_lt2 := cD[0].latitude
        min_lt2 := cD[0].latitude
        max_ln2 := cD[0].longitude
        min_ln2 := cD[0].longitude

        done := make(chan bool) 
        mid:= int(len(cD)/2)

        //This gets the first half of the array
        go func(){                                                                              
            max_lat, min_lat, max_long, min_long = setupv2(cD[:mid], xdm, ydm)
            done<- true
        }()

        go func(){                                                                              
            max_lt2, min_lt2, max_ln2, min_ln2 = setupv2(cD[mid:], xdm, ydm)
            done<- true
        }()
        <-done
        <-done


        if(min_ln2<min_long){
            min_long=min_ln2
        }
        if(max_ln2>max_long){
            max_long=max_ln2
        }
        if(min_lt2<min_lat){
            min_lat=min_lt2
        }
        if(max_lt2>max_lat){
            max_lat=max_lt2
        }

    	return max_lat, min_lat, max_long, min_long
	}
}


func calPop(cD []CensusGroup, xdv float64, ydv float64, minlt float64, minln float64, w int, s int, e int, n int) (int, int){
    if len(cD)<20000 {
        total_pop:=0
        population:=0
        for y:= 0; y<len(cD); y++{
            new_x := int((cD[y].longitude - minln)/xdv)+1
            new_y := int((cD[y].latitude - minlt)/ydv)+1
            
            if ((new_x >= w) && (new_x <= e) && (new_y >= s) &&  (new_y <= n)){
                population = population + cD[y].population 
            }
            // fmt.Println(population)
            total_pop = total_pop + cD[y].population
        }
        return population,total_pop
    } else{
        t1:=0
        p1:=0
        t2:=0
        p2:=0
        mid:= len(cD)/2
        taskcomplete := make(chan bool) 
        go func(){
            p1,t1 = calPop(cD[mid:], xdv, ydv, minlt, minln, w, s, e, n)
            taskcomplete<-true
        }()
        go func(){
            p2, t2 = calPop(cD[:mid], xdv, ydv, minlt, minln, w, s, e, n)
            taskcomplete<-true
        }()
        <-taskcomplete
        <-taskcomplete
        return (p1+p2), (t1+t2)
        
    }
}
func arraySum(arr1 [][]int, arr2 [][]int, start int, end int, xdm int, ydm int){
	r:=0
	c:=0
	if ((end-start)<=(xdm*ydm)/2){
		for y:= start; y<end; y++{
			r = y/xdm
			c = y%xdm
			arr1[r][c]=arr1[r][c] + arr2[r][c]
		}
	}else{
		mid:= (end-start)/2
		taskcomplete := make(chan bool) 
        go func(){
            arraySum(arr1, arr2, start, mid, xdm, ydm)
            taskcomplete<-true
        }()
		arraySum(arr1, arr2, mid, end, xdm, ydm)
		<-taskcomplete
	}
}


func makeGrid(cD []CensusGroup, xdv float64, ydv float64, minlt float64, minln float64, xdm int, ydm int) ([][] int){

	if len(cD)<20000 {
        
        func_grid := make([][]int, ydm)
	    for l:= 0; l<ydm; l++{
	        func_grid[l] = make([]int, xdm)
	    }

	    for y:= 0; y<ydm; y++{
            for x:= 0; x<xdm; x++{
                func_grid[y][x] = 0
            }
        }

        for y:= 0; y<len(cD); y++{
            new_x := int((cD[y].longitude - minln)/xdv)+1
            new_y := int((cD[y].latitude - minlt)/ydv)+1

            if ((new_x-1)<xdm && (new_x-1)>-1 && (new_y-1)>-1 && (new_y-1)<ydm){          //Only value (13,4) is being ignored. I printed it(See readme)
                func_grid[new_y-1][new_x-1] = func_grid[new_y-1][new_x-1] + cD[y].population 
            }
        }
        return func_grid
     } else{
     	func_grid := make([][]int, ydm)
     	for l:= 0; l<ydm; l++{
     		func_grid[l] = make([]int, xdm)
     	}
     	for y:= 0; y<ydm; y++{
     		for x:= 0; x<xdm; x++{
     			func_grid[y][x] = 0
     		}
     	}

     	func_grid2 := make([][]int, ydm)
     	for l:= 0; l<ydm; l++{
     		func_grid2[l] = make([]int, xdm)
     	}
     	for y:= 0; y<ydm; y++{
     		for x:= 0; x<xdm; x++{
     			func_grid2[y][x] = 0
     		}
     	}

     	done := make(chan bool)
     	mid:= len(cD)/2
     	go func(){
     		func_grid= makeGrid(cD[:mid], xdv, ydv, minlt, minln, xdm, ydm)
     		done<-true
     	}()
     	go func(){
     		func_grid2= makeGrid(cD[mid:], xdv, ydv, minlt, minln, xdm, ydm)
     		done<-true
     	}()

     	<-done
     	<-done

		//have to prallelize this summing
     	arraySum(func_grid, func_grid2, 0, (xdm*ydm), xdm, ydm)
     
     	return func_grid
     }
}
func makeGridp5(cD []CensusGroup, xdv float64, ydv float64, minlt float64, minln float64, xdm int, ydm int, muGrid [][]*sync.Mutex, grid [][]int ) ([][] int){

	if len(cD)<20000 {
        for y:= 0; y<len(cD); y++{
            new_x := int((cD[y].longitude - minln)/xdv)+1
            new_y := int((cD[y].latitude - minlt)/ydv)+1

            if ((new_x-1)<xdm && (new_x-1)>-1 && (new_y-1)>-1 && (new_y-1)<ydm){          //Only value (13,4) is being ignored. I printed it(See readme)
            	muGrid[new_y-1][new_x-1].Lock()
                grid[new_y-1][new_x-1] = grid[new_y-1][new_x-1] + cD[y].population 
                muGrid[new_y-1][new_x-1].Unlock()
            }
        }  
        return grid
     } else{
     	
     	done := make(chan bool)
     	mid:= len(cD)/2
     	go func(){
     		grid= makeGridp5(cD[:mid], xdv, ydv, minlt, minln, xdm, ydm, muGrid, grid)
     		done<- true
     	}()
     	go func(){
     		grid= makeGridp5(cD[mid:], xdv, ydv, minlt, minln, xdm, ydm, muGrid, grid)
     		done<-true
     	}()

     	<-done
     	<-done
     	return grid
     }
}


func ParseCensusData(fname string) ([]CensusGroup, error) {
	file, err := os.Open(fname)
    if err != nil {
		return nil, err
    }
    defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	censusData := make([]CensusGroup, 0, len(records))

    for _, rec := range records {
        if len(rec) == 7 {
            population, err1 := strconv.Atoi(rec[4])
            latitude, err2 := strconv.ParseFloat(rec[5], 64)
            longitude, err3 := strconv.ParseFloat(rec[6], 64)
            if err1 == nil && err2 == nil && err3 == nil {
                latpi := latitude * math.Pi / 180
                latitude = math.Log(math.Tan(latpi) + 1 / math.Cos(latpi))
                censusData = append(censusData, CensusGroup{population, latitude, longitude})
            }
        }
    }

	return censusData, nil
}

func main () {
	if len(os.Args) < 4 {
		fmt.Printf("Usage:\nArg 1: file name for input data\nArg 2: number of x-dim buckets\nArg 3: number of y-dim buckets\nArg 4: -v1, -v2, -v3, -v4, -v5, or -v6\n")
		return
	}
	fname, ver := os.Args[1], os.Args[4]
    xdim, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
    ydim, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	censusData, err := ParseCensusData(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
    // fmt.Println(censusData)

    len_rec := len(censusData)
    min_lat:=0.0
    min_long:= 0.0
    max_lat:=0.0
    max_long:=0.0
    x_div:= 0.0
    y_div:=0.0
    grid := make([][]int, ydim)
    for l:= 0; l<ydim; l++{
        grid[l] = make([]int, xdim)
    }
    // Some parts may need no setup code
    switch ver {
    case "-v1":
        max_lat = censusData[0].latitude
        min_lat = censusData[0].latitude
        max_long = censusData[0].longitude
        min_long = censusData[0].longitude
        for x:= 0; x<len_rec; x++{
            if (censusData[x].latitude > max_lat){
                max_lat=censusData[x].latitude 
            }
            if (censusData[x].latitude < min_lat){
                min_lat=censusData[x].latitude 
            }
            if (censusData[x].longitude > max_long){
                max_long=censusData[x].longitude 
            }
            if (censusData[x].longitude < min_long){
                min_long=censusData[x].longitude
            }
        }
        x_div = (max_long - min_long)/ float64(xdim)
        y_div = (max_lat - min_lat)/ float64(ydim)
    case "-v2":
        // int mid = censusData                                                                             
        max_lat, min_lat, max_long, min_long = setupv2(censusData, xdim, ydim)


        x_div = (max_long - min_long)/ float64(xdim)
        y_div = (max_lat - min_lat)/ float64(ydim)

    case "-v3":
        max_lat = censusData[0].latitude
        min_lat = censusData[0].latitude
        max_long = censusData[0].longitude
        min_long = censusData[0].longitude
        for x:= 0; x<len_rec; x++{
            if (censusData[x].latitude > max_lat){
                max_lat=censusData[x].latitude 
            }
            if (censusData[x].latitude < min_lat){
                min_lat=censusData[x].latitude 
            }
            if (censusData[x].longitude > max_long){
                max_long=censusData[x].longitude 
            }
            if (censusData[x].longitude < min_long){
                min_long=censusData[x].longitude
            }
        }
        x_div = (max_long - min_long)/ float64(xdim)
        y_div = (max_lat - min_lat)/ float64(ydim)
        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                grid[y][x] = 0
            }
        }
        
        for y:= 0; y<len(censusData); y++{
            new_x := int((censusData[y].longitude - min_long)/x_div)+1
            new_y := int((censusData[y].latitude - min_lat)/y_div)+1
           
            if ((new_x-1)<xdim && (new_x-1)>-1 && (new_y-1)>-1 && (new_y-1)<ydim){          //Only value (13,4) is being ignored. I printed it(See readme)
                grid[new_y-1][new_x-1] = grid[new_y-1][new_x-1] + censusData[y].population 
            }
        }
        fmt.Println(grid)
        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                if(y>0){
                    if(x>0){
                        grid[y][x] = grid[y][x] + grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1]
                    } else{
                        grid[y][x] = grid[y][x] + grid[y-1][x]
                    }  
                } else{
                    if(x>0){
                        grid[y][x] = grid[y][x] +  grid[y][x-1]
                    } else{
                        grid[y][x] = grid[y][x]  
                    }
                }
            }
        }
       
    case "-v4":
        // YOUR SETUP CODE FOR PART 4
        //corner finding in parallel
    	max_lat, min_lat, max_long, min_long = setupv2(censusData, xdim, ydim)
    	// fmt.Println(max_lat)
    	// fmt.Println(max_long)
    	// fmt.Println(min_lat)

        x_div = (max_long - min_long)/ float64(xdim)
        y_div = (max_lat - min_lat)/ float64(ydim)

        //initializing the grid

        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                grid[y][x] = 0
            }
        }

        grid = makeGrid(censusData, x_div, y_div, min_lat, min_long, xdim, ydim)

        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                if(y>0){
                    if(x>0){
                        grid[y][x] = grid[y][x] + grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1]
                    } else{
                        grid[y][x] = grid[y][x] + grid[y-1][x]
                    }  
                } else{
                    if(x>0){
                        grid[y][x] = grid[y][x] +  grid[y][x-1]
                    } else{
                        grid[y][x] = grid[y][x]  
                    }
                }
            }
        }
        // fmt.Println(grid)


    case "-v5":
        // YOUR SETUP CODE FOR PART 5
    	max_lat, min_lat, max_long, min_long = setupv2(censusData, xdim, ydim)


        x_div = (max_long - min_long)/ float64(xdim)
        y_div = (max_lat - min_lat)/ float64(ydim)

        //initializing the grid
        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                grid[y][x] = 0
            }
        }

        mutexGrid := make([][]*sync.Mutex, ydim)
	    for i:=0; i<ydim; i++{
				mutexGrid[i] = make([]*sync.Mutex, xdim)
			}
		//initializing the mutex
		for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
				var mu sync.Mutex
				mutexGrid[y][x] = &mu
			}
		}
        grid = makeGridp5(censusData, x_div, y_div, min_lat, min_long, xdim, ydim, mutexGrid, grid)

        for y:= 0; y<ydim; y++{
            for x:= 0; x<xdim; x++{
                if(y>0){
                    if(x>0){
                        grid[y][x] = grid[y][x] + grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1]
                    } else{
                        grid[y][x] = grid[y][x] + grid[y-1][x]
                    }  
                } else{
                    if(x>0){
                        grid[y][x] = grid[y][x] +  grid[y][x-1]
                    } else{
                        grid[y][x] = grid[y][x]  
                    }
                }
            }
        }
    case "-v6":
        // YOUR SETUP CODE FOR PART 6
    default:
        fmt.Println("Invalid version argument")
        return
    }

    for {
        var west, south, east, north int
        n, err := fmt.Scanln(&west, &south, &east, &north)
        if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
            break
        }

        var population int
        var percentage float64
        switch ver {
        case "-v1":

            total_pop:= 0
     
            for y:= 0; y<len(censusData); y++{
                new_x := int((censusData[y].longitude - min_long)/x_div)+1
                new_y := int((censusData[y].latitude - min_lat)/y_div)+1
                
                if ((new_x >= (west)) && (new_x <= (east)) && (new_y >= (south)) &&  (new_y <= (north))){
                    population = population + censusData[y].population
                }
                total_pop = total_pop + censusData[y].population
            }
            percentage = (float64(population) / float64(total_pop))*100

        case "-v2":

            
            total_pop:= 0
            population, total_pop = calPop(censusData, x_div, y_div, min_lat, min_long, west, south, east, north)
           
            percentage = (float64(population) / float64(total_pop))*100

        case "-v3":
            
            var a,b,c,d int
            a= grid[north-1][east-1]
            if ((south-2)>-1 && (east-1)>-1){
                b = grid[south -2][east - 1]
            }else{
                b = 0
            }
            if ((west-2)>-1 && (north-1)>-1){
                c = grid[north -1][west -2]
            }else{
                c = 0
            }
            if ((south-2)>-1 && (west-2)>-1){
                d = grid[south -2][west - 2]
            }else{
                d = 0
            }
            population = a - b - c + d
            // population = grid[7][5] - grid[1][5] - grid[7][1] + grid[1][1]
            percentage = (float64(population) / float64(grid[ydim-1][xdim-1]))*100
          
        case "-v4":
            // YOUR QUERY CODE FOR PART 4
            var a,b,c,d int
            a= grid[north-1][east-1]
            if ((south-2)>-1 && (east-1)>-1){
                b = grid[south -2][east - 1]
            }else{
                b = 0
            }
            if ((west-2)>-1 && (north-1)>-1){
                c = grid[north -1][west -2]
            }else{
                c = 0
            }
            if ((south-2)>-1 && (west-2)>-1){
                d = grid[south -2][west - 2]
            }else{
                d = 0
            }
            population = a - b - c + d
            // population = grid[7][5] - grid[1][5] - grid[7][1] + grid[1][1]
            percentage = (float64(population) / float64(grid[ydim-1][xdim-1]))*100
        case "-v5":
            // YOUR QUERY CODE FOR PART 5
            var a,b,c,d int
            a= grid[north-1][east-1]
            if ((south-2)>-1 && (east-1)>-1){
                b = grid[south -2][east - 1]
            }else{
                b = 0
            }
            if ((west-2)>-1 && (north-1)>-1){
                c = grid[north -1][west -2]
            }else{
                c = 0
            }
            if ((south-2)>-1 && (west-2)>-1){
                d = grid[south -2][west - 2]
            }else{
                d = 0
            }
            population = a - b - c + d
            // population = grid[7][5] - grid[1][5] - grid[7][1] + grid[1][1]
            percentage = (float64(population) / float64(grid[ydim-1][xdim-1]))*100
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}
