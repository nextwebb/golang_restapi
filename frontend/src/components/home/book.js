import React from "react";
import axios from "axios";

let endpoint = "http://localhost:8080";



class Book extends React.Component{
    constructor(){
        super();
        this.state = {
            ID: "",
            Isbn: "",
            Title: "",
            Author: ""
        }
    }
 

    render(){
          return (

         <div className="container">
         <p className="h3"> BOOK API</p>

    <form>
       <div className="header">

       </div>
<div className="form-group">
 <label for="exampleInputEmail1">ID</label>
 <input type="text" className="form-control" id="ID" aria-describedby="emailHelp" placeholder="Book ID"/>
 
</div>
<div className="form-group">
<label for="exampleInputEmail1">Isbn</label>
 <input type="text" className="form-control" id="event"  placeholder="ISBN"/>
</div>
<div className="form-group">
 <label for="exampleInputPassword1">Title</label>
 <input type="text" className="form-control" id="participants" placeholder="Book Title"/>
</div>
<div className="form-group">
 <label for="exampleInputPassword1">Author</label>
 <input type="text" className="form-control" id="participants" placeholder="Author Name"/>
</div>

<button type="submit" className="btn btn-success mr-5">Post Book</button>
<button type="submit" className="btn btn-primary mr-5">Edit Book</button>
<button type="submit" className="btn btn-danger mr-5">Delete Book</button>

    </form>
      </div>
          )}
}

export default Book;