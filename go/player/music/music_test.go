package music

import "testing"
import "fmt"

func TestA(t *testing.T){
    l := NewLib()
    l.Add("Deng Li Jun", "Yue liang dai biao wo de xin")
    l.Add("Beyonce", "Halo")
    //fmt.Println(l.Len())
    l.Show()

    l.Delete(1)
    l.Show()
    
    l.Play(0)
    fmt.Println("==over==")
}