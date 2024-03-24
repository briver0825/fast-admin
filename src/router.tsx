import { createBrowserRouter, Navigate } from "react-router-dom";

import DashboardLayout from "./layouts/dashboard";
import EmptyLayout from "./layouts/empty";

import NotFound from "./components/NotFound";

import LoginPage from "./pages/login";
import UserPage from "./pages/user";
import ShopPage from './pages/shop'
import ShopCategoryPage from './pages/shop/category'

const router = createBrowserRouter([
    {
        path: '/',
        element: <DashboardLayout />,
        children: [
            {
                path: '',
                element: <Navigate replace={true} to={'/user/list'} />
            },
            {
                path: 'user',
                element: <EmptyLayout />,
                children: [
                    {
                        path: 'list',
                        element: <UserPage />,
                    }
                ]
            },
            {
                path: 'shop',
                element: <EmptyLayout />,
                children: [
                    {
                        path: 'list',
                        element: <ShopPage />,
                    },
                    {
                        path: 'category',
                        element: <ShopCategoryPage />
                    }
                ]
            }
        ],
    },
    {
        path: '/login',
        element: <LoginPage />
    },
    {
        path: "*",
        element: <NotFound />
    }
])

export default router