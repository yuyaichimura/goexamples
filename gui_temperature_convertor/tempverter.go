package main

import (
    "github.com/andlabs/ui"
    "strconv"
    "fmt"
)

/* A global reference to ui.Window */
var window ui.Window

/* Initializes the gui window */
func initGUI() {

    /* Initializes components
     * Textfields, buttons, labels
     */
    tf_celsius := ui.NewTextField()
    tf_fahrenheit := ui.NewTextField()
    b_c2f := ui.NewButton(">>")
    b_f2c := ui.NewButton("<<")
    l_c := ui.NewLabel("C")
    l_f := ui.NewLabel("F")

    /* Creates Stacks to align the components the way 
     * I want them to align. Could have used grid
     * layouts or some other layouts but I decided
     * to use vertical and horizontal stacks 
     */
    stack_c := ui.NewHorizontalStack(tf_celsius, l_c)
    stack_c2 := ui.NewVerticalStack(stack_c, b_c2f)

    stack_f := ui.NewHorizontalStack(tf_fahrenheit, l_f)
    stack_f2 := ui.NewVerticalStack(stack_f, b_f2c)

    stack := ui.NewSimpleGrid(2,
        stack_c2,
        stack_f2)

    stack.SetPadded(true)

    /* Sets the title and size of the window */
    window = ui.NewWindow("Temperature Calculator", 370, 70, stack)
    window.SetMargined(true)
    window.OnClosing(func() bool {
        ui.Stop()
        return true
    })

    /* Sets the OnClicked functions for the buttons
     * using anonymous functions */
    b_c2f.OnClicked(func(){

        var res string
        var val_f float64

        val_c, err := strconv.ParseFloat(tf_celsius.Text(), 64)
        if err != nil {
            res = "ERR!"
        } else {
            val_f = ((val_c * 1.8) + 32)
            if val_f == float64(int(val_f)) {
                res = fmt.Sprintf("%.0f", val_f)
            } else {
                res = fmt.Sprintf("%.2f", val_f)
            }
        }

        tf_fahrenheit.SetText(res)
    })

    b_f2c.OnClicked(func(){

        var res string
        var val_c float64

        val_f, err := strconv.ParseFloat(tf_fahrenheit.Text(), 64)
        if err != nil {
            res = "ERR!"
        } else {
            val_c = (val_f - 32) / 1.8
            if val_c == float64(int(val_c)) {
                res = fmt.Sprintf("%.0f", val_c)
            } else {
                res = fmt.Sprintf("%.2f", val_c)
            }
        }

        tf_celsius.SetText(res)
    })

    /* Finally shows the window*/
    window.Show()
}

/* Main function 
 * Calls a function to initialize the gui on another thread
 */
func main() {
    go ui.Do(initGUI)
    err := ui.Go()
    if err != nil {
        panic(err)
    }
}