import {
  createRequestHandler,
  defaultRenderHandler,
} from '@tanstack/react-router/ssr/server'
import { createRouter } from './router'
import { StrictMode } from 'react'
import { renderToString } from 'react-dom/server'
import { ChakraProvider } from '@chakra-ui/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { App } from './main'

// export default async function render({ request }: { request: Request }) {
//   const handler = createRequestHandler({ request, createRouter })

//   return await handler(defaultRenderHandler)
// }

// export function render(_url: string) {
//   const html = renderToString(
//     <StrictMode>
//       <App />
//     </StrictMode>,
//   )
//   return { html }
// }

const queryClient = new QueryClient()

export function render(_url: string) {
  const html = renderToString(
    <StrictMode>
      <ChakraProvider>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </ChakraProvider>
    </StrictMode>,
  )
  return { html }
}
