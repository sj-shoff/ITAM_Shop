import React, { useState } from "react"
import { Button, Card, Input, Spacer } from "@nextui-org/react"

export function LoginPage() {
    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")

    const handleLogin = () => {
        console.log("Logging in with:", { email, password })
    }

    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "100vh",
                backgroundColor: "#f5f5f5",
            }}
        >
            <Card style={{ padding: "20px", width: "100%", maxWidth: "400px" }}>
                <h3 style={{ textAlign: "center", marginBottom: "20px" }}>
                    Login
                </h3>
                <Input
                    type='email'
                    fullWidth
                    placeholder='Email'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <Spacer y={1.5} />
                <Input
                    type='password'
                    fullWidth
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <Spacer y={1.5} />
                <Button color='primary' onClick={handleLogin} fullWidth>
                    Login
                </Button>
            </Card>
        </div>
    )
}
