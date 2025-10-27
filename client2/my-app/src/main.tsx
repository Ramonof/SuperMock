import { StrictMode } from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from '@tanstack/react-router'

import './styles.css'
import reportWebVitals from './reportWebVitals.ts'

import { ChakraProvider } from '@chakra-ui/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { useAuth, AuthProvider } from './context/auth.tsx'
import { createRouter } from './router.tsx'

// export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8080/api/v1" : "/api/v1";
export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8080/api/v1" : "http://localhost:8080/api/v1";

const queryClient = new QueryClient()

// Create a new router instance
// const router = createRouter({
//   routeTree,
//   context: {
//     auth: undefined!,
//   },
  // defaultPreload: 'intent',
  // scrollRestoration: true,
  // defaultStructuralSharing: true,
  // defaultPreloadStaleTime: 0,
// })

const router = createRouter()

function InnerApp() {
  const auth = useAuth()
  return <RouterProvider router={router} context={{ auth }} />
}

export function App() {
  return (
    <AuthProvider>
      <InnerApp />
    </AuthProvider>
  )
}

// Render the app
// const rootElement = document.getElementById('app')
// if (rootElement && !rootElement.innerHTML) {
//   const root = ReactDOM.createRoot(rootElement)
//   root.render(
//     <StrictMode>
//       <ChakraProvider>
//         <QueryClientProvider client={queryClient}>
//           <App />
//         </QueryClientProvider>
//       </ChakraProvider>
//     </StrictMode>,
//   )
// }

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
