package main
import "fmt"

func main() {
    var n,q int
    fmt.Scanf("%d %d",&n,&q)
    seq := make([][]int,n)
    qseq := make([][]int,q)
    //var qseq [q][3]int
    for i:=0;i<q;i++{
        qseq[i]= make([]int,3) 
        for j:=0;j<3;j++{
            fmt.Scanf("%d",&qseq[i][j])
        }
    }
    for i:=0;i<n;i++{
        seq[i] = make([]int,0)
    }
    la:=0
    for i:=0;i<q;i++{
        switch qseq[i][0]{
        case 1:
            //fmt.Println(qseq)
            ind := (qseq[i][1]^la)%n
            seq[ind] =append(seq[ind],qseq[i][2])
            
            //fmt.Println(seq)
        case 2:
            //fmt.Println(2)
            //fmt.Println(qseq)
            //fmt.Println(seq)
            ind1 := (qseq[i][1]^la)%n
            ind2:= qseq[i][2]%len(seq[ind1])
            la = seq[ind1][ind2]
                fmt.Println(la)
           
        }
    }
    
}
