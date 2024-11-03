import backend_config from './config'
let base_url = backend_config.backend_url
   base_url = "http://" + base_url + "/get"

export default async function () {
    return await fetch(base_url, {
        method: "get",
        headers: {
            "Access-Control-Allow-Origin": "*"
        }
    }).then((res)=>{
        return res.json()
	}).then((res)=>{
        return res
		})
}