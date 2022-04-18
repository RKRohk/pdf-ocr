import { useEffect, useState } from "react"

interface LoadingProps {
    id: string
}
const Loading:React.FC<LoadingProps> = ({id}) => {

    const [message,setMessage] = useState<string>("")

    const initWS = async () => {
        await new Promise(resolve => setTimeout(resolve,1000))
        const ws = new WebSocket(`wss://${window.location.hostname}/ocr/ws/`+id)
        ws.onopen = (ev) => {
            console.log("socket opened")
        }

        ws.onmessage = (ev) => {
            console.log("event received ",ev)
            setMessage(ev.data)
        }
    }
    useEffect(() => {
       initWS()
        
    },[id])

    return <>
        <p className="text-center">
            {message}
        </p>
    </>
}

export default Loading