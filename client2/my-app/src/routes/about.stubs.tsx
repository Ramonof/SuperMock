import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/about/stubs')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/about/stubs"!</div>
}
