import {
    LockOutlined,
    UserOutlined,
} from '@ant-design/icons';
import {
    LoginFormPage,
    ProFormText,
} from '@ant-design/pro-components';

import bgVideo from '@/assets/login-bg.mp4'
import { useNavigate } from 'react-router-dom';
import { message } from 'antd';

export default () => {
    const navigate = useNavigate()

    return (
        <div className='h-screen'>
            <LoginFormPage
                title="FastAdmin"
                subTitle="让开发变得简单"
                backgroundVideoUrl={bgVideo}
                onFinish={async (_) => {
                    sessionStorage.setItem('token', "token")
                    navigate('/')
                    message.success('登录成功')
                }}
            >
                <ProFormText
                    name="username"
                    fieldProps={{
                        size: 'large',
                        prefix: <UserOutlined className={'prefixIcon'} />,
                    }}
                    placeholder={'用户名'}
                    rules={[
                        {
                            required: true,
                            message: '请输入用户名!',
                        },
                    ]}
                />
                <ProFormText.Password
                    name="password"
                    fieldProps={{
                        size: 'large',
                        prefix: <LockOutlined className={'prefixIcon'} />,
                    }}
                    placeholder={'密码'}
                    rules={[
                        {
                            required: true,
                            message: '请输入密码！',
                        },
                    ]}
                />
            </LoginFormPage>
        </div >
    );
};