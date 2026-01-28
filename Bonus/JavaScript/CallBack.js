console.log("Hello Friends");

function callBack(callBack)
{
    setTimeout(() => {
        var api_resposne = "Hello";
        callBack(api_resposne);
        
    },3000);
    
}
function myCallFun(message)
{
    console.log(message);

}

callBack(myCallFun);
console.log("I am Working without callback");