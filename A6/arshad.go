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
func set_1 (censusData []CensusGroup) (float64,float64,float64,float64) {
    min_long := censusData[0].latitude
        min_lat := censusData[0].latitude
        max_long := censusData[0].longitude
        max_lat := censusData[0].latitude

        for i:=1;i<len(censusData);i++ {
            if(censusData[i].latitude<min_lat){
                min_lat=censusData[i].latitude
            }
            if(censusData[i].latitude>max_lat){
                max_lat=censusData[i].latitude
            }
            if(censusData[i].longitude<min_long){
                min_long=censusData[i].longitude
            }
            if(censusData[i].longitude>max_long){
                max_long=censusData[i].longitude
            }
        }
        // fmt.Println("min_lat:", min_lat)
        // fmt.Println("max_lat:", max_lat)
        // fmt.Println("min_long", min_long)
        // fmt.Println("max_long",max_long)
        return min_long,max_long,min_lat,max_lat
}
func set_2 (censusData []CensusGroup) (float64,float64,float64,float64) {
    cut_off := 1000
    if(len(censusData)<cut_off){
        return set_1(censusData)
    } else{
        done := make(chan bool)
        var min_long_l, max_long_l, min_lat_l, max_lat_l float64
        var min_long_r, max_long_r, min_lat_r, max_lat_r float64
        go func() {
            min_long_l, max_long_l, min_lat_l, max_lat_l = set_2(censusData[:(len(censusData)/2)])
            done<-true
        }()
        min_long_r, max_long_r, min_lat_r, max_lat_r = set_2(censusData[(len(censusData)/2):])
        <-done
        min_lo := math.Min(min_long_r,min_long_l)
        min_la := math.Min(min_lat_l,min_lat_r)
        max_la := math.Max(max_lat_r,max_lat_l)
        max_lo := math.Max(max_long_r,max_long_l)
        return min_lo,max_lo,min_la,max_la
    }
}
func first_grid_lock(censusData []CensusGroup, arr_2d [][]int, muGrid [][]*sync.Mutex, xdim int, ydim int, min_lat float64,max_lat float64,min_long float64,max_long float64) ([][]int,int) {
    tot :=0
    y_div := (max_lat-min_lat)/float64(ydim)
    x_div := (max_long-min_long)/float64(xdim)
    for j:=0;j<len(censusData);j++ {
        x_cor := math.Ceil((censusData[j].longitude - min_long)/x_div)
        y_cor := math.Ceil((censusData[j].latitude - min_lat)/y_div)
        if(x_cor>0){
        x_cor--;
        }
        if(y_cor>0) {
           y_cor--;
        }
        muGrid[int(x_cor)][int(y_cor)].Lock()
        arr_2d[int(x_cor)][int(y_cor)] = arr_2d[int(x_cor)][int(y_cor)] + censusData[j].population
        muGrid[int(x_cor)][int(y_cor)].Unlock()
        tot = tot + censusData[j].population
    }
    return arr_2d,tot
}
func first_grid(censusData []CensusGroup, xdim int, ydim int, min_lat float64,max_lat float64,min_long float64,max_long float64) ([][]int,int) {
    arr_2d := make([][]int, xdim)
    for i:=0;i<xdim;i++ {
        arr_2d[i] = make([]int, ydim)
    }
    tot :=0
    y_div := (max_lat-min_lat)/float64(ydim)
    x_div := (max_long-min_long)/float64(xdim)
    for j:=0;j<len(censusData);j++ {
        x_cor := math.Ceil((censusData[j].longitude - min_long)/x_div)
        y_cor := math.Ceil((censusData[j].latitude - min_lat)/y_div)
        if(x_cor>0){
        x_cor--;
        }
        if(y_cor>0) {
           y_cor--;
        }
        arr_2d[int(x_cor)][int(y_cor)] = arr_2d[int(x_cor)][int(y_cor)] + censusData[j].population
        tot = tot + censusData[j].population
    }
    return arr_2d,tot
}
func rem_set_3(arr_2d_left [][]int, xdim int, ydim int) ([][]int){
    for m:=1;m<ydim;m++ {
        arr_2d_left[0][m] = arr_2d_left[0][m] + arr_2d_left[0][m-1] 
    }
    for n:=1;n<xdim;n++ {
        arr_2d_left[n][0] = arr_2d_left[n][0] + arr_2d_left[n-1][0]
    }
    for a:=1;a<xdim;a++ {
        for b:=1;b<ydim;b++ { 
                arr_2d_left[a][b] = arr_2d_left[a][b] + arr_2d_left[a-1][b] + arr_2d_left[a][b-1] - arr_2d_left[a-1][b-1]
    }
    }
    return arr_2d_left

}
func add_two(arr_2d_left [][]int, arr [][]int, arr_2d_right [][]int,row int) ([][]int){
	cutoff:=3
	if(len(arr_2d_left)<cutoff){
		for i:=row;i<(len(arr_2d_left)+row);i++{
			for j:=0;j<len(arr_2d_left[0]);j++ {
				arr[i][j] = arr_2d_left[i-row][j] + arr_2d_right[i-row][j]
			}
		}
		return arr
	} else {
		done := make(chan bool)
		go func () {
			arr = add_two(arr_2d_left[(len(arr_2d_left)/2):],arr,arr_2d_right[(len(arr_2d_right)/2):],(row+(len(arr_2d_left)/2)))
			done<-true
		}()
		arr = add_two(arr_2d_left[:(len(arr_2d_left)/2)],arr,arr_2d_right[:(len(arr_2d_right)/2)],row)
		<-done
		return arr
	}

}
func set_4(censusData []CensusGroup, xdim int, ydim int,min_lat float64,max_lat float64,min_long float64,max_long float64) ([][]int,int) {
    cut_off := 10000
    done := make(chan bool)
        var tot_l,tot_r int
        arr_2d_left := make([][]int, xdim)
    for c:=0;c<xdim;c++ {
        arr_2d_left[c] = make([]int, ydim)
    }
    arr_2d_right := make([][]int, xdim)
    for d:=0;d<xdim;d++ {
        arr_2d_right[d] = make([]int, ydim)
    }
    if(len(censusData)<cut_off){
        return first_grid(censusData,xdim,ydim,min_lat,max_lat,min_long,max_long)
    } else {
    go func() {
        arr_2d_left,tot_l = set_4(censusData[(len(censusData)/2):],xdim,ydim,min_lat,max_lat,min_long,max_long)
        done<-true
    }()
    arr_2d_right,tot_r = set_4(censusData[:(len(censusData)/2)],xdim,ydim,min_lat,max_lat,min_long,max_long)
    <-done
    arr_2d_left = add_two(arr_2d_left,arr_2d_left,arr_2d_right,0)
    return arr_2d_left,(tot_r+tot_l)
}
}
func set_5(censusData []CensusGroup, arr_2d [][]int, muGrid [][]*sync.Mutex, xdim int, ydim int,min_lat float64,max_lat float64,min_long float64,max_long float64) ([][]int,int) {
    cut_off := 10000
    var tot_l,tot_r int
    done := make(chan bool)
    if(len(censusData)<cut_off){
        return first_grid_lock(censusData,arr_2d, muGrid,xdim,ydim,min_lat,max_lat,min_long,max_long)
    } else {
    go func() {
        arr_2d,tot_l = set_5(censusData[(len(censusData)/2):],arr_2d, muGrid,xdim,ydim,min_lat,max_lat,min_long,max_long)
        done<-true
    }()
    arr_2d,tot_r = set_5(censusData[:(len(censusData)/2)],arr_2d, muGrid,xdim,ydim,min_lat,max_lat,min_long,max_long)
    <-done
    // fmt.Println(tot_r+tot_l)
    // fmt.Println(arr_2d)
    return arr_2d,(tot_l+tot_r)
}
}
func set_3(censusData []CensusGroup, xdim int, ydim int) ([][]int,int) {
    var min_lat,max_lat,min_long,max_long float64
    var tot int
    min_long,max_long,min_lat,max_lat = set_1(censusData)

    arr_2d := make([][]int, xdim)
    for i:=0;i<xdim;i++ {
        arr_2d[i] = make([]int, ydim)
    }
    y_div := (max_lat-min_lat)/float64(ydim)
    x_div := (max_long-min_long)/float64(xdim)
    for j:=0;j<len(censusData);j++ {
        x_cor := math.Ceil((censusData[j].longitude - min_long)/x_div)
        y_cor := math.Ceil((censusData[j].latitude - min_lat)/y_div)
        if(x_cor>0){
        x_cor--;
        }
        if(y_cor>0) {
           y_cor--;
        }
        // fmt.Println(int(x_cor))
        // fmt.Println(int(y_cor))
        arr_2d[int(x_cor)][int(y_cor)] = arr_2d[int(x_cor)][int(y_cor)] + censusData[j].population
        tot = tot + censusData[j].population
    }
    for m:=1;m<ydim;m++ {
        arr_2d[0][m] = arr_2d[0][m] + arr_2d[0][m-1] 
    }
    for n:=1;n<xdim;n++ {
        arr_2d[n][0] = arr_2d[n][0] + arr_2d[n-1][0]
    }
    for a:=1;a<xdim;a++ {
        for b:=1;b<ydim;b++ { 
                arr_2d[a][b] = arr_2d[a][b] + arr_2d[a-1][b] + arr_2d[a][b-1] - arr_2d[a-1][b-1]
    }
    }
    return arr_2d,tot
}

func query_1 (censusData []CensusGroup, west int, south int, east int, north int, max_long float64, min_long float64, min_lat float64, max_lat float64, xdim int, ydim int) (int,float64,int) {
    tot := 0
    var pop int
    var per float64
            y_div := (max_lat-min_lat)/float64(ydim)
            //fmt.Println(y_div)
            x_div := (max_long-min_long)/float64(xdim)
            //fmt.Println(x_div)
            for j:=0;j<len(censusData);j++ {

                x_cor := math.Ceil((censusData[j].longitude - min_long)/x_div)
                y_cor := math.Ceil((censusData[j].latitude - min_lat)/y_div)
                if(x_cor==0){
                    x_cor++
                }
                if(y_cor==0) {
                    y_cor++;
                }

                if((int(x_cor)>=west && int(x_cor)<=east) && (int(y_cor)>=south && int(y_cor)<=north)) {
                    pop = pop + censusData[j].population
                }
            
                tot = tot + censusData[j].population
            }
                per = (float64(pop)/float64(tot))*100
                return pop, per, tot              
}
func query_2(censusData []CensusGroup, west int, south int, east int, north int, max_long float64, min_long float64, min_lat float64, max_lat float64, xdim int, ydim int) (int,float64,int) {
    cut_off := 10000
    var left_pop int
    var tot_left int
    var tot_right int
    var right_pop int
    if(len(censusData)<cut_off){
        return query_1(censusData,west,south,east,north,max_long,min_long,min_lat,max_lat,xdim,ydim)
    } else {
        done := make(chan bool)
        go func() {
            left_pop,_, tot_left = query_2(censusData[:(len(censusData)/2)],west,south,east,north,max_long,min_long,min_lat,max_lat,xdim,ydim)
            done<-true
        }()
        right_pop,_, tot_right = query_2(censusData[(len(censusData)/2):],west,south,east,north,max_long,min_long,min_lat,max_lat,xdim,ydim)
        <-done
        population := left_pop + right_pop
        tot := tot_right+tot_left
        percentage := (float64(population)/float64(tot))*100
        return population,percentage,tot
    }
}

func query_3(arr_2d [][]int,west int, south int, east int, north int, total int) (int,float64) {
var pop int
var per float64
	//fmt.Println(arr_2d)
	west--;
	south--;
	north--;
	east--;
	var se, nw, sw int
	if (west-1>=0) {
	se = arr_2d[west-1][north]
} else {
	se = 0
}
if(south-1>=0) {
	nw = arr_2d[east][south-1]
} else {
	nw = 0
}
if(south-1>=0 && west-1>=0){
	sw = arr_2d[west-1][south-1]
} else {
	sw = 0
}
pop = arr_2d[east][north] - se - nw + sw
per = (float64(pop)/float64(total))*100


return pop,per
}

func ParseCensusData (fname string) ([]CensusGroup, error) {
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
    var min_lat,max_lat,min_long,max_long float64
    var total int
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
    arr_2d := make([][]int, xdim)
    for i:=0;i<xdim;i++ {
        arr_2d[i] = make([]int, ydim)
    }
    // Some parts may need no setup code
    switch ver {
    case "-v1":
        min_long,max_long,min_lat,max_lat = set_1(censusData)
    case "-v2":
        min_long,max_long,min_lat,max_lat = set_2(censusData)
        fmt.Println("min_lat:", min_lat)
        fmt.Println("max_lat:", max_lat)
        fmt.Println("min_long", min_long)
        fmt.Println("max_long",max_long)
    case "-v3":
        // YOUR SETUP CODE FOR PART 3
        arr_2d,total = set_3(censusData,xdim, ydim)
        
    case "-v4":
        // YOUR SETUP CODE FOR PART 4
        min_long,max_long,min_lat,max_lat = set_2(censusData)
        arr_2d,total = set_4(censusData,xdim,ydim,min_lat,max_lat,min_long,max_long)
        arr_2d = rem_set_3(arr_2d,xdim,ydim)
        fmt.Println("SETUP DONE")
    case "-v5":
    	muGrid := make([][]*sync.Mutex, xdim)
     for i:=0; i<xdim; i++ {
			muGrid[i] = make([]*sync.Mutex, ydim)
		}
		for i:=0;i<xdim;i++ {
			for j:=0; j<ydim; j++ {
				var mu sync.Mutex
				muGrid[i][j] = &mu
			}
		}
    	min_long,max_long,min_lat,max_lat = set_2(censusData)
        arr_2d,total = set_5(censusData,arr_2d,muGrid,xdim,ydim,min_lat,max_lat,min_long,max_long)
        fmt.Println(arr_2d)
        arr_2d = rem_set_3(arr_2d,xdim,ydim)
        fmt.Println("SETUP DONE")
        // YOUR SETUP CODE FOR PART 5
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
            population, percentage, _ = query_1(censusData,west,south,east,north,max_long,min_long,min_lat,max_lat,xdim,ydim)
        case "-v2":
            population, percentage, _ = query_2(censusData,west,south,east,north,max_long,min_long,min_lat,max_lat,xdim,ydim)
        case "-v3":
            population, percentage = query_3(arr_2d,west,south,east,north,total)
        case "-v4":
            // YOUR QUERY CODE FOR PART 4
            population, percentage = query_3(arr_2d,west,south,east,north,total)
        case "-v5":
        	population, percentage = query_3(arr_2d,west,south,east,north,total)
            // YOUR QUERY CODE FOR PART 5
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}