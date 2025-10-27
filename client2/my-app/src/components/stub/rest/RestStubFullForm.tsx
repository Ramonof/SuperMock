import { BASE_URL } from "@/main";
import { Flex, Input, Button, Spinner, Stack, Text, HStack, Select, Radio, RadioGroup, useColorModeValue, createMultiStyleConfigHelpers, extendTheme, useColorMode } from "@chakra-ui/react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useEffect, useRef, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import type { RestStub } from "./RestStubList";
import apiClient from "@/utils/request";
import {EditorView, basicSetup} from "codemirror"
import {json, jsonParseLinter} from "@codemirror/lang-json"
import {esLint, javascript} from "@codemirror/lang-javascript"
import {oneDark} from "@codemirror/theme-one-dark"
import CodeMirror, { useCodeMirror } from '@uiw/react-codemirror';
import React from "react";
import { linter } from "@codemirror/lint";
import { radioAnatomy } from '@chakra-ui/anatomy'
import * as eslint from "eslint-linter-browserify";

import {syntaxTree} from "@codemirror/language"
import type {Diagnostic} from "@codemirror/lint"
import getColor from "@/utils/color";

const bracketsLinter = linter(view => {
  let diagnostics: Diagnostic[] = []
  let a = 0
  let b = 0
  syntaxTree(view.state).cursor().iterate(node => {
    b = node.to
    if (node.name == "{") {
        a = a + 1
    }
    if (node.name == "}") {
        a = a - 1
    }
    if (a < 0) {
        diagnostics.push({
            from: node.from,
            to: node.to,
            severity: "warning",
            message: "Unopened bracket",
        })
    }
  })
  if (a > 0) {
            diagnostics.push({
            from: b,
            to: b,
            severity: "warning",
            message: "Unclosed bracket",
        })
  }
  return diagnostics
})

// Define the extensions outside the component for the best performance.
// If you need dynamic extensions, use React.useMemo to minimize reference changes
// which cause costly re-renders.
const extensionsJson = [json(), linter(jsonParseLinter()), javascript()];
const extensionsJs = [javascript(), bracketsLinter];

const RestStubFullForm = ({ ProjectId, StubId, restStubs, setRestStubs }: { ProjectId: string, StubId: string, restStubs: RestStub[], setRestStubs: React.Dispatch<React.SetStateAction<RestStub[]>> }) => {
    const color = useColorModeValue("gray.800", "yellow.100")
    const { colorMode, toggleColorMode } = useColorMode();
    
    const [newRestStubName, setNewRestStubName] = useState("");
    const [newRestStubPath, setNewRestStubPath] = useState("");
    const [newRestStubMethod, setNewRestStubMethod] = useState("");
    const [newRestStubType, setNewRestStubType] = useState("json")
    const [newRestStubResponseBody, setNewRestStubResponseBody] = useState("");

    const [extensions, setExtensions] = useState(extensionsJson)

    const onTypeChange = React.useCallback((val: React.SetStateAction<string>) => {
        setNewRestStubType(val)
        if (val === "json") {
            setExtensions(extensionsJson)
        } else {
            setExtensions(extensionsJs)
        }
    }, []);

    const onChange = React.useCallback((val: React.SetStateAction<string>, viewUpdate: any) => {
        setNewRestStubResponseBody(val)
    }, []);

    const editor = useRef(null);
    const { setContainer } = useCodeMirror({
        container: editor.current,
        extensions,
        value: newRestStubResponseBody,
        onChange: onChange,
        height: "200px",
        theme: colorMode === "light" ? "light" : "dark"
    });

    useEffect(() => {
        if (editor.current) {
            setContainer(editor.current);
        }
    }, [editor.current]);

    const navigate = useNavigate();

    const queryClient = useQueryClient();

    const { data: stubData, isLoading } = useQuery<RestStub>({
		queryKey: ["getStubById" + StubId],
		queryFn: async () => {
			try {
				const response = await apiClient({
					url: BASE_URL + "/projects/" + ProjectId + "/stub/" + StubId, 
					method: 'get'
				});

				const data = response.data
                if (newRestStubName == "") {
                    setNewRestStubName(data.name);
                    setNewRestStubPath(data.path);
                    setNewRestStubMethod(data.method);
                    setNewRestStubResponseBody(data.response_body)
                    setNewRestStubType(data.type)
                }

				return data || null;
			} catch (error) {
				console.log(error);
			}
		},
	});

    const { mutate: updateRestStub, isPending: isCreating } = useMutation({
        mutationKey: ["updateRestStub"],
        mutationFn: async (e: React.FormEvent) => {
            e.preventDefault();
            try {
                const res = await apiClient({
					url: BASE_URL + "/projects/" + ProjectId + "/stub/" + StubId, 
					method: 'PUT',
                    headers: {
                        "Content-Type": "application/json",
                    },
                    data: JSON.stringify({ name: newRestStubName, path: newRestStubPath, method: newRestStubMethod, type: newRestStubType, response_body: newRestStubResponseBody }),
				});

                setRestStubs(restStubs.map(item => (item.id == res.data.id ? res.data : item)));
                
                // setNewRestStubName("");
                // setNewRestStubPath("");
                // setNewRestStubMethod("")
                // setNewRestStubResponseBody("");
                // navigate({
                //     to: '/project/$projectId/rest/stubs/$stubId',
                //     params: { projectId: ProjectId, stubId: StubId },
                // })
                return res;
            } catch (error: any) {
                throw new Error(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["getRestStubs"] });
        },
        onError: (error: any) => {
            alert(error.message);
        },
    });

    const [arrId, setArrId] = useState(1)
    const inputArr = [
        {
            type: "text",
            id: arrId,
            value: ""
        }
    ];
    const [arr, setArr] = useState(inputArr);
    const addInput = () => {
        setArrId(arrId+1)
        setArr(s => {
        return [
            ...s,
            {
            type: "text",
            id: arrId,
            value: ""
            }
        ];
        });
    };

    const handleChange = (e: { preventDefault: () => void; target: { id: any; value: string; }; }) => {
        e.preventDefault();

        const index = e.target.id;
        setArr(s => {
        const newArr = s.slice();
        newArr[index].value = e.target.value;

        return newArr;
        });
    };

    return (
        <form onSubmit={updateRestStub}>
            <Stack gap={2}>
                <Text color={getColor()}  fontSize='xl'>
                    Request
                </Text>
                <Text color={getColor()}>
                    Name
                </Text>
                <Input
                    type='text'
                    value={newRestStubName}
                    onChange={(e) => setNewRestStubName(e.target.value)}
                    // ref={(input) => input && input.focus()}
                    textFillColor={getColor()}
                />
                <Text color={getColor()}>
                    Path
                </Text>
                <HStack>
                    <Select value={newRestStubMethod} onChange={(e) => setNewRestStubMethod(e.target.value)} color={getColor()} width={'fit-content'}>
                        <option value='ANY'>ANY</option>
                        <option value='GET'>GET</option>
                        <option value='POST'>POST</option>
                        <option value='PUT'>PUT</option>
                        <option value='PATCH'>PATCH</option>
                        <option value='DELETE'>DELETE</option>
                        <option value='HEAD'>HEAD</option>
                        <option value='OPTIONS'>OPTIONS</option>
                    </Select>
                    <Input
                        type='text'
                        value={newRestStubPath}
                        onChange={(e) => setNewRestStubPath(e.target.value)}
                        // ref={(input) => input && input.focus()}
                        textFillColor={getColor()}
                    />
                </HStack>
                <Text color={getColor()}  fontSize='xl'>
                    Response
                </Text>
                <button onClick={addInput}><Text color={getColor()}>+ Header</Text></button>
                {arr.map((item, i) => {
                    return (
                    <HStack>
                    <Input
                        onChange={handleChange}
                        value={item.value}
                        id={i}
                        type={item.type}
                        size="40"
                        textFillColor={getColor()}
                    />
                    <Input
                        onChange={handleChange}
                        value={item.value}
                        id={i}
                        type={item.type}
                        size="40"
                        textFillColor={getColor()}
                    />
                    </HStack>
                    );
                })}
                <Text color={getColor()} >Body</Text>
                <RadioGroup onChange={onTypeChange} value={newRestStubType}>
                    <Stack direction='row'>
                        <Radio value='json' textColor={getColor()} colorScheme='green'>
                            <Text color={color}>json</Text>
                        </Radio>
                        <Radio textColor={color} value='javascript' colorScheme='red'>
                            <Text color={color}>javascript(WIP)</Text>
                        </Radio>
                        <Radio textColor={color} value='goroovy' colorScheme='green'>
                            <Text color={color}>goroovy</Text>
                        </Radio>
                    </Stack>
                </RadioGroup>
                {/* <CodeMirror value={newRestStubResponseBody} height="200px" extensions={[json(), linter(jsonParseLinter()), javascript()]} onChange={onChange} /> */}
                <div ref={editor} />
                <Button
                    mx={2}
                    type='submit'
                    _active={{
                        transform: "scale(.97)",
                    }}
                >
                    {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />} Save
                </Button>
            </Stack>
        </form>
    );
};
export default RestStubFullForm;