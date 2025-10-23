import { hydrateRoot } from 'react-dom/client'
import { RouterClient } from '@tanstack/react-router/ssr/client'
import { createRouter } from './router'
import { StrictMode } from 'react'
import { ChakraProvider } from '@chakra-ui/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { App } from './main'

// const router = createRouter()

// hydrateRoot(document, <RouterClient router={router} />)

// hydrateRoot(
//   document.getElementById('root') as HTMLElement,
//   <StrictMode>
//     <App />
//   </StrictMode>,
// )

const queryClient = new QueryClient()

hydrateRoot(
  document.getElementById('app') as HTMLElement,
  <StrictMode>
      <ChakraProvider>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </ChakraProvider>
    </StrictMode>,
)