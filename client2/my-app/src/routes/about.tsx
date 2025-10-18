import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/about')({
  component: About,
})

function About() {
  const { auth } = Route.useRouteContext()
  return <div className="p-2">Hello from About! <strong>{auth.user?.user}</strong>!</div>
}