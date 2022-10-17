package gosh

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "runtime"
)

// `ShellOutput` returns the output of shell command, and any errors.
func ShellOutput(command string, environmentVariables []string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("bash", "-c", command)
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  cmd.Env = environmentVariables

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

func ShellOutputWithDir(command, dir string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("bash", "-c", command)
  cmd.Dir = dir
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

// `ShellCommand` executes the shell command.
func ShellCommand(command string, environmentVariables []string) {
  err, out, errout := ShellOutput(command, environmentVariables)

  if err != nil {
    log.Printf("error: %v\n", err)
    fmt.Print(errout)
  }

  fmt.Print(out)
}

// `PowershellOutput` returns the output of powershell command, and any errors.
func PowershellOutput(command string, environmentVariables []string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("powershell.exe", command)
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  cmd.Env = environmentVariables

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

func PowershellOutputWithDir(command, dir string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("powershell.exe", command)
  cmd.Dir = dir
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

// `PowershellCommand` executes the powershell command.
func PowershellCommand(command string, environmentVariables []string) {
  err, out, errout := PowershellOutput(command, environmentVariables)

  if err != nil {
    log.Printf("error: %v\n", err)
    fmt.Print(errout)
  }

  fmt.Print(out)
}

// `Exec` just exectes the command
func Exec(shell, command string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command(shell, "-c", command)

  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

// `Run` executes the same command for shell and powershell
func Run(cmd string, environmentVariables []string) {
  err, out, errout := ShellOutput("", environmentVariables)

  if runtime.GOOS == "windows" {
    err, out, errout = PowershellOutput(cmd, environmentVariables)
  } else {
    err, out, errout = ShellOutput(cmd, environmentVariables)
  }

  if err != nil {
    log.Printf("error: %v\n", err)
    fmt.Print(errout)
  }

  fmt.Print(out)
}

// `RunOutput` returns the output of the shared command for shell and powershell
func RunOutput(command string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("")

  if runtime.GOOS == "windows" {
    cmd = exec.Command("powershell.exe", command)
  } else {
    cmd = exec.Command("bash", "-c", command)
  }

  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}

// `RunMulti` executes a command for shell and a command for powershell
func RunMulti(unixCmd string, winCmd string, environmentVariables []string) {
  err, out, errout := ShellOutput("", environmentVariables)

  if runtime.GOOS == "windows" {
    err, out, errout = PowershellOutput(winCmd, environmentVariables)
  } else {
    err, out, errout = ShellOutput(unixCmd, environmentVariables)
  }

  if err != nil {
    log.Printf("error: %v\n", err)
    fmt.Print(errout)
  }

  fmt.Print(out)
}

// `RunMultiOutput` returns the output of the shell command and the powershell command
func RunMultiOutput(unixCmd string, winCmd string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer

  cmd := exec.Command("")

  if runtime.GOOS == "windows" {
    cmd = exec.Command("powershell.exe", winCmd)
  } else {
    cmd = exec.Command("bash", "-c", unixCmd)
  }

  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  err := cmd.Run()

  return err, stdout.String(), stderr.String()
}
