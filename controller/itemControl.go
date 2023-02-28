package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (c *Collection) ItemPush(name string, price, rating float64) {
	c.itemIterator++
	item := Item{c.itemIterator, name, price, rating}
	c.Items = append(c.Items, item)
	c.ItemSaveData()
	fmt.Println("Item:", item, "was created successfully.")
}

func (c *Collection) SearchItemsByName(name string) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.ItemName == name {
			result = append(result, item)
		}
	}
	return result
}

func (c *Collection) FilterItemsByPrice(price float64) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.Price <= price {
			result = append(result, item)
		}
	}

	return result
}

func (c *Collection) FilterItemsByRating(rating float64) []Item {
	var result []Item
	for _, item := range c.Items {
		if item.Rating <= rating {
			result = append(result, item)
		}
	}
	return result
}

func (c *Collection) SetRating(rating float64, id int) {
	i := 0
	for ; i < len(c.Items); i++ {
		if c.Items[i].Id == id {
			c.Items[i].Rating = rating
			c.ItemSaveData()
			return
		}
	}
	fmt.Println("item was not found.")
	return
}

func (c *Collection) ItemTakeData() {
	file, err := os.Open("item.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Split(line, ",")
		n, err := strconv.Atoi(i[0])
		if err != nil {
			fmt.Println(err)
		}
		flo, err := strconv.ParseFloat(i[2], 64)
		if err != nil {
			fmt.Println(err)
		}
		flot, err := strconv.ParseFloat(i[3], 64)
		if err != nil {
			fmt.Println(err)
		}
		item := Item{n, i[1], flo, flot}
		c.Items = append(c.Items, item)
		c.itemIterator = n
	}
}

func (c *Collection) ItemSaveData() {
	if err := os.Truncate("item.txt", 0); err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("item.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, item := range c.Items {
		if _, err := file.WriteString(strconv.Itoa(item.Id) + "," + item.ItemName +
			"," + strconv.FormatFloat(item.Price, 'g', 5, 64) +
			"," + strconv.FormatFloat(item.Rating, 'g', 5, 64) + "\n"); err != nil {
			fmt.Println(err)
		}
	}
}
