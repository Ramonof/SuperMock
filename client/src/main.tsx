import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { Provider } from './components/ui/provider.tsx'
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, createBrowserRouter, RouterProvider } from 'react-router'

const queryClient = new QueryClient();
const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/project/:ProjectId",
    // loader: async ({ params }) => {
    //   let team = await fetchTeam(params.teamId);
    //   return { name: team.name };
    // },
    element: <App />,
  }
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
    <BrowserRouter>
    <Provider>
    <App />
    {/* <RouterProvider router={router} /> */}
    {/* <App /> */}
    </Provider>
    </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>,
)
