import { Avatar, Button, Dropdown, Menu, message, Space } from "antd"
import { MenuProps } from "antd/lib/menu";
import { useEffect, useMemo, useState } from "react";

import { AccountBookOutlined, DownOutlined, MenuFoldOutlined, MenuUnfoldOutlined, ShopOutlined, ShoppingOutlined, UserOutlined } from '@ant-design/icons'
import { Navigate, useLocation, useNavigate, useOutlet } from "react-router-dom";

import NotFound from "@/components/NotFound";

type ISidebarProps = {
    collapsed: boolean
}

const Sidebar: React.FC<ISidebarProps> = ({ collapsed }) => {
    const location = useLocation()
    const [current, setCurrent] = useState(location.pathname);
    const navigate = useNavigate()

    const items: MenuProps['items'] = [
        {
            label: '用户管理',
            key: '/user/list',
            icon: <UserOutlined />
        },
        {
            label: '商品管理',
            key: '/shop',
            icon: <ShopOutlined />,
            children: [
                {
                    label: '商品列表',
                    key: '/shop/list',
                    icon: <ShoppingOutlined />
                },
                {
                    label: '商品分类',
                    key: '/shop/category',
                    icon: <AccountBookOutlined />
                },
            ],
        }
    ];

    useEffect(() => {
        navigate(current)
    }, [current])

    return (
        <div className=" bg-white transition-all delay-100 overflow-hidden" style={{ borderRight: '1px solid #E5E7ED', width: collapsed ? 50 : 200 }}>
            <div className="text-center py-4 font-bold whitespace-nowrap">
                {collapsed ? 'F' : 'FastAdmin 管理后台'}
            </div>
            <Menu selectedKeys={[current]} mode="inline" items={items} onClick={e => setCurrent(e.key)} inlineCollapsed={collapsed} />
        </div>
    )
}

function UserDropdown() {
    const navigate = useNavigate()
    return (
        <Dropdown menu={{
            items: [
                {
                    key: 1,
                    label: '退出登录',
                    onClick: () => {
                        sessionStorage.removeItem('token')
                        navigate('/login')
                        message.success("退出成功")
                    }
                }
            ]
        }}>
            <div className="cursor-pointer hover:bg-gray-100 h-[50px] flex items-center px-2">
                <Space>
                    <Avatar icon={<UserOutlined />} />
                    <span className="text-gray-500">
                        Admin <DownOutlined />
                    </span>
                </Space>
            </div>
        </Dropdown>
    )
}

type IHeaderProps = {
    toggleCollapsed: () => any
    collapsed: boolean
}

const Header: React.FC<IHeaderProps> = ({ collapsed, toggleCollapsed }) => {
    return (
        <div className="bg-white flex items-center justify-between px-4" style={{ borderBottom: '1px solid #E5E7ED' }}>
            <Button type="primary" onClick={toggleCollapsed}>
                {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            </Button>
            <UserDropdown />
        </div>
    )
}


function Content() {
    const outlet = useOutlet()
    return (
        <div className="m-4">
            {outlet ?? <NotFound />}
        </div>
    )
}

function Layout() {
    const [collapsed, setCollapsed] = useState(false);

    return (
        <div className="flex h-screen bg-[#F6F6F6]">
            <Sidebar collapsed={collapsed} />
            <div className="flex-1">
                <Header collapsed={collapsed} toggleCollapsed={() => setCollapsed(!collapsed)} />
                <Content />
            </div>
        </div>
    )
}

export default () => {
    const location = useLocation()
    const token = sessionStorage.getItem("token")

    const layout = useMemo(() => <Layout />, [])

    return (
        !!token ? layout : <Navigate replace to={`/login?redirect=${encodeURIComponent(location.pathname + location.search)}`} />
    )
}