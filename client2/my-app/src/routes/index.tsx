import { Flex, useColorModeValue, Stack, Heading, Container, Text, VStack, Box, HStack } from '@chakra-ui/react'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: Index,
})

function Index() {
  return (
        <Stack spacing={10}>
          <Container maxWidth='120ch'>
            <Heading size='2xl' textAlign={'center'}>
              Совершенная платформа симуляции API
            </Heading>
          </Container>
          <Container maxWidth='120ch'>
            <Text fontSize='lg' textAlign={'center'}>
              Запускается везде где запускается ваш код. AI-ready с нативной поддержкой MCP.
            </Text>
          </Container>
          <HStack>
            <Container maxWidth='120ch'>
              <Box bg={useColorModeValue("gray.400", "gray.700")}>
                <Heading size='xl' textAlign={'center'}>
                  Получите
                </Heading>
                <Text>
                  - Тысячи часов в скорости
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более надежная скорость параллельной разработки
                </Text>
              </Box>
            </Container>
            <Container maxWidth='120ch'>
              <Box bg={useColorModeValue("gray.400", "gray.700")}>
                <Heading size='xl' textAlign={'center'}>
                  Сократите
                </Heading>
                <Text>
                  - Тысячи часов в скорости
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более надежная скорость параллельной разработки
                </Text>
              </Box>
            </Container>
          </HStack>
          <Container maxWidth='120ch'>
            <HStack>
              <Container maxWidth='-webkit-max-content'>
                <Heading size='xl'>
                  Точный
                </Heading>
                <Text>
                  Моделируемые API, отражающие реальность
                </Text>
                <Text>
                  - Соответствие запросов, ошибки, задержка ответа, динамические ответы
                </Text>
                <Text>
                  - Динамические состояния во время сеансов тестирования
                </Text>
                <Text>
                  - Динамические состояния сеансов тестирования и мощные сценарии
                </Text>
              </Container>
              <Container maxWidth='-webkit-max-content'>
                <Box boxShadow={'0 0 40px #0550ff40'}>
                  <img src='/main_page.png' alt='logo' width={'100%'} height={'100%'}/>
                </Box>
              </Container>
            </HStack>
          </Container>
          <Container maxWidth='120ch'>
            <HStack margin={'auto'}>
              <Container maxWidth='-webkit-max-content'>
                <Box boxShadow={'0 0 40px #0550ff40'}>
                  <img src='/spec_gen.png' alt='logo' width={'100%'} height={'100%'}/>
                </Box>
              </Container>
              <Container maxWidth='-webkit-max-content'>
                <Heading size='xl'>
                  Гибкий
                </Heading>
                <Text>
                  Получите максимум от OpenAPI
                </Text>
                <Text>
                  - Храните файлы OpenAPI в Git, отправляйте и извлекайте изменения в свои макеты по мере необходимости
                </Text>
                <Text>
                  - Создание прототипов с использованием макетов для быстрого проектирования новых API
                </Text>
                <Text>
                  - Автоматическая генерация спецификаций OpenAPI
                </Text>
              </Container>
            </HStack>
          </Container>
        </Stack>
  )
}