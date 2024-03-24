import { useOutlet } from "react-router-dom"
import NotFound from "@/components/NotFound"

export default () => {
    const outlet = useOutlet()
    return <>{outlet ?? <NotFound />}</>
}