package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func main() {
	m, err := MD5All2(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for p := range m {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s \n ", m[path], path)
	}
}

func MD5All(file string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)
	m := make(map[string][md5.Size]byte)
	//filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
	//	if err != nil{
	//		return err
	//	}
	//	if !info.Mode().IsRegular(){
	//		//not ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
	//		return nil
	//	}
	//	data, err := ioutil.ReadFile(path)
	//	if err != nil{
	//		return err
	//	}
	//	m[path] = md5.Sum(data)
	//	return nil
	//})
	c, errc := sumFiles(done, file)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

// parallel
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{
					path: path,
					sum:  md5.Sum(data),
					err:  err,
				}:
				case <-done:
				}
				wg.Done()
			}()
			select {
			case <-done:
				return errors.New("walk cancelled !")
			default:
				return nil
			}
		})
		go func() {
			wg.Wait()
			close(c)
		}()
		errc <- err
	}()
	return c, errc
}

func walkFiles(done chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)// 这里很重要
	//errc := make(chan error)
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")

			}
			return nil
		})
	}()
	return paths, errc
}

func digester(done chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{
			path: path,
			sum:  md5.Sum(data),
			err:  err,
		}:
		case <-done:
			return
		}
	}
}

func MD5All2(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)
	paths, errc := walkFiles(done, root)
	c := make(chan result)
	var wg sync.WaitGroup
	const number = 2
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if e := <-errc; e != nil {
		return nil, e
	}
	return m, nil
}
