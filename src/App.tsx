import { ConfigProvider } from "antd"

import router from "./router"
import { RouterProvider } from "react-router-dom"
import { ProConfigProvider } from "@ant-design/pro-components"

function App() {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: '#4A5DFF',
          borderRadius: 4,
        },
        components: {
          Menu: {
            itemHeight: 46,
            itemMarginBlock: 0,
            itemMarginInline: 0,
            itemBorderRadius: 0,
            collapsedWidth: 50
          }
        },
      }}>
      <ProConfigProvider hashed={false}>
        <RouterProvider router={router} />
      </ProConfigProvider>
    </ConfigProvider>
  )
}

export default App