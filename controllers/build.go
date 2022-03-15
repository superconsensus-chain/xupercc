package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/superconsensus-chain/xupercc/conf"
	log "github.com/superconsensus-chain/xupercc/utils"
)

//ccfile: 带后缀的是文件，不带的是包
//output: xx.wasm
func BuildCC(ccfile string) error {
	/*
	 * 因为从虚拟机clone的xupercc调用contract_deploy总是部署失败，且摸索之后发现xdev编译单个.cc文件总是报错，遂改
	 * 通常报错信息如下，换了好几个版本的xdev都无解
	   make: .Makefile: No such file or directory
	   make: *** No rule to make target '.Makefile'. Stop.
	   Error: exit status 2
	   exit status 2
	   详细测试记录见文档[xxx]
	*/
	// 不再使用xxx.cc文件直接build，将文件复制到一个新的目录后再xdev build
	ccdir := strings.Split(ccfile, ".cc")
	// 生成合约相关文件夹
	os.Mkdir(conf.Code.CodePath, 0755)
	os.Mkdir(conf.Code.WasmPath, 0755)

	// cc文件构建成的临时工程结构目录（绝对路径，方便中间有错误返回时使用相对路径没删除成功）
	buildDir := filepath.Join(conf.Code.CodePath, ccdir[0])
	buildDir, _ = filepath.Abs(buildDir)
	defer func() {
		os.RemoveAll(buildDir)
	}()

	os.MkdirAll(filepath.Join(buildDir, "src"), 0755)
	// 生成xdev.toml文件
	writeErr := ioutil.WriteFile(filepath.Join(buildDir, "xdev.toml"), []byte(`[package]
	name = "main"
	`), 0644)
	if writeErr != nil {
		log.Println("write [xdev.toml] file err.", writeErr)
		return writeErr
	}

	// 将.cc源文件复制到临时工程的src下
	_, err := exec.Command("cp", conf.Code.CodePath+ccfile, filepath.Join(buildDir, "src")).Output()
	if err != nil {
		log.Println("cp file err. origin file=", conf.Code.CodePath+ccfile, "destination=", buildDir+"/src", "err=", err)
		return err
	}

	// 设置xdev编译环境
	xdevPath, err := filepath.Abs("xdev")
	if err != nil {
		log.Println("获取xdev绝对路径失败", err)
		return err
	}
	err = os.Setenv("XDEV_ROOT", xdevPath)
	if err != nil {
		log.Println("设置XDEV_ROOT环境失败", err)
		return err
	}

	// 切换目录，等效cd命令
	os.Chdir(buildDir)
	// xdev编译工程
	cmd := exec.Command("../../xdev/xdev", "build")

	log.Println(cmd.String())

	output, err := cmd.Output()
	if err != nil {
		log.Printf("c合约编译错误，可能是docker服务没有启动; $XDEV_ROOT环境: %s", os.Getenv("XDEV_ROOT"))
		os.Chdir("../../")
		return err
	}

	if !strings.Contains(string(output), "LD wasm") {
		log.Println("c合约编译失败", string(output))
		os.Chdir("../../")
		return fmt.Errorf("build desc: cmd: %s output: %s", cmd.String(), string(output))
	}
	// 移动编译好的文件后回到项目根目录
	exec.Command("mv", ccdir[0]+".wasm", "../../"+conf.Code.WasmPath).Output()
	os.Chdir("../../")

	return nil
}

//output: xx.go.native
func BuildGo(gofile string) error {
	// 目前测试了4个go版本（1.13.4/1.14.5/1.15.9/1.16.13）只有1.13build出来的wasm能部署，所以直接使用go build，取代原先编译wasm
	nativefile := gofile + ".native" //go的合约加上后缀
	in := filepath.Join(conf.Code.CodePath, gofile)
	out := filepath.Join(conf.Code.WasmPath, nativefile)
	cmd := exec.Command("go", "build", "-o", out, in)

	log.Println(cmd.String())

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(output) != "" {
		return fmt.Errorf("build desc: cmd: %s output: %s", cmd.String(), string(output))
	}

	return nil
}

func BuildSol(solfile string) error {
	in := filepath.Join(conf.Code.CodePath, solfile+".sol")

	// 规范约束，如果sol文件中不包含solfile这样的contract的话返回错误，但是可能使用注释绕过判断
	codeBytes, err := ioutil.ReadFile(in)
	if err != nil {
		return fmt.Errorf("读取sol源文件失败")
	}
	if !strings.Contains(string(codeBytes), solfile) {
		return fmt.Errorf("合约名(contract_name)应与源代码中声明的主合约名一致")
	}

	// 编译 solc --bin --abi Counter.sol -o .
	cmd := exec.Command("./xdev/solc", "--bin", "--abi", in, "-o", filepath.Join(conf.Code.WasmPath, solfile), "--overwrite")
	output, err := cmd.Output()
	if err != nil {
		log.Println("sol编译失败", string(output), "error=", err)
		return err
	}
	/*// solc编译出来的bin跟abi文件名是根据sol文件里声明的contract xxx决定，所以需要额外用正则匹配出来
	reg, err := regexp.Compile("contract\\s+.+{")
	if err != nil {
		return err
	}
	slice := reg.FindAllString(string(codeBytes), -1) // -1 表示列出所有符合的match
	// 匹配0及以上个空格，并用###替换，后续再使用###切割
	spaceReg, err := regexp.Compile("\\s+")
	if err != nil {
		return err
	}
	replace := spaceReg.ReplaceAllString(slice[0], "###")
	split := strings.Split(replace, "###")
	err = exec.Command("mv", filepath.Join(conf.Code.WasmPath, solfile, split[1]+".bin"), filepath.Join(conf.Code.WasmPath, solfile, solfile+".bin")).Run()
	if err != nil {
		return err
	}
	err = exec.Command("mv", filepath.Join(conf.Code.WasmPath, solfile, split[1]+".abi"), filepath.Join(conf.Code.WasmPath, solfile, solfile+".abi")).Run()*/
	return nil
}
