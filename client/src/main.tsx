import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { Provider } from './components/ui/provider.tsx'
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, createBrowserRouter, RouterProvider } from 'react-router'
import {AuthProvider} from './context/AuthProvider.tsx'
import { routeTree } from './routeTree.gen'
import { createRouter } from '@tanstack/react-router'

const queryClient = new QueryClient();

// Create a new router instance
const router = createRouter({ routeTree })

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
    <AuthProvider>
    <Provider>
    <RouterProvider router={router} />
    <App />
    {/* <RouterProvider router={router} /> */}
    {/* <App /> */}
    </Provider>
    </AuthProvider>
    </QueryClientProvider>
  </StrictMode>,
)
