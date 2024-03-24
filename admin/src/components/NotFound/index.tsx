import { Button, Result } from "antd"
import { useNavigate } from "react-router-dom"

const NotFound: React.FC = () => {
    const navigate = useNavigate()
    return < Result
        status="404"
        title="404"
        subTitle="抱歉，这个页面不存在"
        extra={<Button type="primary" onClick={() => navigate(-1)}> 返回上一页</Button>}
    />
}

export default NotFound