function promptMessage(){
    let toast = function(){
        console.log("clicked button and got toast")
    }

    let success = function(){
        console.log("clicked button and got success")
    }

    return {
        toast: toast,
        success: success,
    }
}