const express = require('express')
const bodyParser = require('body-parser');
const notifier = require('node-notifier')
const app = express();
const port = process.env.PORT || 9000;
const path = require('path')

app.use(bodyParser.json());


app.get("/health",(req,res)=> res.status(200).send());
app.post("/notify",(req,res)=>{
    notify(req.body, reply => res.send(reply))
})
app.listen(port,()=> console.log("Server is runing"));

const notify = ({title,message}, cb) => {
    notifier.notify(
        {
            title: title || "Some Title",
            message: message || "Some Message",
            sound:true,
            icon:path.join(__dirname,"Tipu.jpg"),
            wait:true,
            reply:true,
            closeLabel:"Completed",
            timeout:15
        },
        (err, response, metadata) => {
            cb(metadata)
        } 
    )
}