package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
	"strconv"
	"time"
)

func error_handler(err error){
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func correctHour(s string) string {

	// init vars
	rebuild_bool := false
	result := s

	// extract hour
	text_array := strings.Split(s, ",")
	arrival_time := text_array[1]
	departure_time := text_array[2]

	arrival_time_array := strings.Split(arrival_time, ":")
	departure_time_array := strings.Split(departure_time, ":")

	arrival_hour_str := arrival_time_array[0]
	departure_hour_str := departure_time_array[0]

	arrival_hour, err := strconv.ParseInt(arrival_hour_str, 10, 64)
	error_handler(err)
	departure_hour, err := strconv.ParseInt(departure_hour_str, 10, 64)
	error_handler(err)

	if arrival_hour >= 24{
		arrival_time_array[0] = correct_single_hour(arrival_hour)
		arrival_time =  strings.Join(arrival_time_array, ":")
		text_array[1] = arrival_time
		rebuild_bool = true
	}

	if departure_hour >= 24{
		departure_time_array[0] = correct_single_hour(departure_hour)
		departure_time = strings.Join(departure_time_array, ":")
		text_array[2] = departure_time
		rebuild_bool = true
	}

	if rebuild_bool {
		result = strings.Join(text_array, ",")
	}

	return result
}

func correct_single_hour(int_num int64) string {

	result := int_num%24
	result_str := fmt.Sprintf("%02d", result)

	return result_str
}

func main() {

	// start
	start := time.Now()

	// read from input
	input_filename := os.Args[1]
	output_filename := os.Args[2]
	// open file
	input_file, err := os.Open(input_filename)
	defer input_file.Close()
	error_handler(err)

	//fmt.Println(input_filename)
	//fmt.Println(output_filename)

	// create output file
	output_file, err := os.Create(output_filename)
	defer output_file.Close()
	error_handler(err)

	// open reader and writer
	reader := bufio.NewReader(input_file)
	writer := bufio.NewWriter(output_file)

	// read and write the first line
	line_of_text, err := reader.ReadString('\n')
	_ , err = writer.WriteString(line_of_text)
	writer.Flush()

	// for loop
	for {
		counter := 1
		line_of_text, err := reader.ReadString('\n')
		//fmt.Println(line_of_text)
		// break if end of file
		if err == io.EOF {
			break
		}

		// Process the line
		output_string := correctHour(line_of_text)
		//fmt.Println(output_string)

		// write to output file
		_ , err = writer.WriteString(output_string)

		if counter >= 1000 {
			writer.Flush()
			counter++
		} else {
			counter = 1
		}

		error_handler(err)
	}

	// flush unwritten lines
	writer.Flush()

	// end
	elapsed := time.Since(start)
	fmt.Printf("Time Check %s", elapsed)
}