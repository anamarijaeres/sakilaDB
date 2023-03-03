import React, { useState } from "react";
import { Card,Grid,Input ,Button,Icon, GridColumn,Reveal,Dropdown,Label,Popup} from "semantic-ui-react";
import axios from "axios";

import "../../index.css"



function ActorCard({ data ,func,vis}){
    const [open,setOpen]=useState(true)
    const [film, setFilms] = useState([]);
    const [openMod,setOpenMod]=useState(false)




    const handleClickOpen = (id) => {
        setOpen(!open)

        if(open){
        console.log("card")

        console.log(open)

        axios
        .get("http://localhost:8080/actors/filmsById/"+id)
        .then((result) => {
            console.log(open)
            console.log(result.data);
            setFilms(result.data);
            func(open,result.data)
            
        })
        .catch((error) => console.log(error));

        }else{
            func(open,[]) 
        }

        console.log(vis)
        console.log(vis)


    }

    const handleDelete= (id)=>{
        const tit=data.Title
        console.log(tit)
        console.log('Delete clicked')
        const url="http://localhost:8080/actors/"+id
        console.log(url)
        axios
        .delete(url)
        .then((result) => {
            console.log(result.data);
            console.log(result.status)
            alert("You have deleted actor "+data.FirstName+" "+data.LastName)
            
        })
        .catch((error) =>{
        alert("You cannot delete actor "+data.FirstName+" "+data.LastName+", since you haven't created it!")
        console.log(error)});
    

    }





    return(
        <Card key={data.ActorId}textAlign="left">
        <Card.Content header={data.FirstName} />
        <Card.Content description={data.LastName} />
        
        <Card.Content extra textAlign="left">
            <Button className=".btn" basic color='blue' content={vis ?  'Show films':(open?'Show films':'Close') } onClick={()=>handleClickOpen(data.ActorId)} / >
           
            
            <div className="delete_btn">
                

                
                <Popup
                    trigger={<Icon name="delete" onClick={()=>handleDelete(data.ActorId)} />}
                    content='Delete actor :('
                    inverted
                    position="right center"
                    />
               
            </div>
            
        </Card.Content>
    </Card>

    )
}
export default ActorCard