import { createFileRoute } from '@tanstack/react-router'
import { Text, useColorModeValue } from '@chakra-ui/react'
import { Link } from '@tanstack/react-router'

export const Route = createFileRoute('/about/overview')({
  component: RouteComponent,
})

function RouteComponent() {
    const color = useColorModeValue("gray.800", "yellow.100")
  return <div>
    <Text color={color}>Легко проектируйте, импортируйте или записывайте любые новые виртуализированные API.</Text>
    <Text color={color}>Управление командой с помощью SSO и RBAC.</Text>
    <Text color={color}>Журналирование активности для мониторинга использования сервисов.</Text>
    <Text color={color}>Внедрение ошибок, хаоса и других форм поведения «на черный день».</Text>
    <img src='/design_chaos_1.png' alt='logo' width={400} height={400} />
    <Text color={color}>Создание синтетических тестовых данных или импорт данных из CSV-файла или базы данных.</Text>
  </div>
}
