package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"

    "github.com/atotto/clipboard"
)

type Item struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

const dataFile = "cliboard.json"

func loadItems() ([]Item, error) {
    file, err := os.ReadFile(dataFile)
    if err != nil {
        if os.IsNotExist(err) {
            return []Item{}, nil
        }
        return nil, err
    }

    var items []Item
    err = json.Unmarshal(file, &items)
    return items, err
}

func saveItems(items []Item) error {
    data, err := json.MarshalIndent(items, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(dataFile, data, 0644)
}

func printItems(items []Item) {
    fmt.Println("+----+--------+-------------------------+")
    fmt.Println("| ID | Name   | Value                   |")
    fmt.Println("+----+--------+-------------------------+")
    for i, item := range items {
        fmt.Printf("| %-2d | %-6s | %-23s |\n", i, item.Name, item.Value)
    }
    fmt.Println("+----+--------+-------------------------+")
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cliboard [add|list|copy|remove]")
        return
    }

    command := os.Args[1]

    items, err := loadItems()
    if err != nil {
        fmt.Println("Error loading data:", err)
        return
    }

    switch command {
    case "add":
        if len(os.Args) < 4 {
            fmt.Println("Usage: cliboard add \"Name\" \"Value\"")
            return
        }
        name := os.Args[2]
        value := os.Args[3]
        items = append(items, Item{Name: name, Value: value})
        if err := saveItems(items); err != nil {
            fmt.Println("Error saving item:", err)
            return
        }
        fmt.Println("Item added.")

    case "list":
        printItems(items)

    case "copy":
        if len(os.Args) < 3 {
            fmt.Println("Usage: cliboard copy <ID>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil || id < 0 || id >= len(items) {
            fmt.Println("Invalid ID.")
            return
        }
        if err := clipboard.WriteAll(items[id].Value); err != nil {
            fmt.Println("Clipboard error:", err)
            return
        }
        fmt.Printf("Copied: %s\n", items[id].Value)

    case "remove":
        if len(os.Args) < 3 {
            fmt.Println("Usage: cliboard remove <ID>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil || id < 0 || id >= len(items) {
            fmt.Println("Invalid ID.")
            return
        }
        // Remove the item by slicing
        items = append(items[:id], items[id+1:]...)
        if err := saveItems(items); err != nil {
            fmt.Println("Error saving items:", err)
            return
        }
        fmt.Printf("Removed item %d.\n", id)

    default:
        fmt.Println("Unknown command:", command)
    }
}

