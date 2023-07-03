package util

import (
	"coderunning/internal/models"
	"coderunning/pkg/constants"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

var TaskNum = make(chan struct{}, 100)

// TODO: 调用Docker
func consume(ctx context.Context, bot models.Bot) error {
	// 创建docker-client
	// client.WithAPIVersionNegotiation() 自动选择API版本
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	fmt.Println("docker连接成功")

	// 创建要运行的go文件
	fileName := createCodeFile(bot)

	timeout := int((2 * time.Second).Seconds())
	config := &container.Config{
		Image:       "276895edf967",
		Cmd:         []string{"go", "run", "/go/src/" + fileName},
		StopTimeout: &timeout,
	}

	hostConfig := &container.HostConfig{
		Binds: []string{constants.VolumeName},
	}
	//  创建容器
	containerId, err := CreateContainer(cli, config, hostConfig, nil, "")
	if err != nil {
		return err
	}

	// 启动容器
	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(containerId)

	// 等待容器内运行结束
	statusCh, errCh := cli.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	// 打印容器内部日志，获得输出
	out, err := cli.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	//stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	resp := make([]byte, 1024)
	n, err := out.Read(resp)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	step := string(resp[:n])
	// TODO: 调用game的SetNextStep

	// 取出一个任务
	<-TaskNum
	fmt.Println(step)

	err = cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("删除容器%s成功\n", containerId)

	return nil
}

func createCodeFile(bot models.Bot) string {
	fileName := fmt.Sprintf("%d.go", len(TaskNum))
	file, _ := os.Create("/home/lzuser/kob/coderunning/util/" + fileName)
	defer file.Close()
	//fmt.Println(file.Name())

	file.WriteString(bot.BotCode)

	return fileName
}

func CreateContainer(
	cli *client.Client,
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	containerName string) (containerId string, err error) {

	ctx := context.Background()

	//创建容器
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, nil, containerName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return resp.ID, nil
}
