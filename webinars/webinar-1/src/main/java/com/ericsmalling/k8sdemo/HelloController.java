package com.ericsmalling.k8sdemo;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
public class HelloController {
    @RequestMapping("/")
    public String index() {
        return "Hello from Kube Academy, "+ System.getProperty("user.name")+" !\n";
    }
}
