// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.GetProjectArgs;
import com.pulumiverse.harbor.inputs.GetProjectPlainArgs;
import com.pulumiverse.harbor.inputs.GetRegistryArgs;
import com.pulumiverse.harbor.inputs.GetRegistryPlainArgs;
import com.pulumiverse.harbor.outputs.GetProjectResult;
import com.pulumiverse.harbor.outputs.GetRegistryResult;
import java.util.concurrent.CompletableFuture;

public final class HarborFunctions {
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetProjectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;library&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, main.applyValue(getProjectResult -&gt; getProjectResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectResult> getProject(GetProjectArgs args) {
        return getProject(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetProjectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;library&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, main.applyValue(getProjectResult -&gt; getProjectResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args) {
        return getProjectPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetProjectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;library&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, main.applyValue(getProjectResult -&gt; getProjectResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectResult> getProject(GetProjectArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("harbor:index/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetProjectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getProject(GetProjectArgs.builder()
     *             .name(&#34;library&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;projectId&#34;, main.applyValue(getProjectResult -&gt; getProjectResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("harbor:index/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetRegistryArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getRegistry(GetRegistryArgs.builder()
     *             .name(&#34;test_docker_harbor&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;harborRegistryId&#34;, main.applyValue(getRegistryResult -&gt; getRegistryResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetRegistryResult> getRegistry(GetRegistryArgs args) {
        return getRegistry(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetRegistryArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getRegistry(GetRegistryArgs.builder()
     *             .name(&#34;test_docker_harbor&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;harborRegistryId&#34;, main.applyValue(getRegistryResult -&gt; getRegistryResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetRegistryResult> getRegistryPlain(GetRegistryPlainArgs args) {
        return getRegistryPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetRegistryArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getRegistry(GetRegistryArgs.builder()
     *             .name(&#34;test_docker_harbor&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;harborRegistryId&#34;, main.applyValue(getRegistryResult -&gt; getRegistryResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetRegistryResult> getRegistry(GetRegistryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("harbor:index/getRegistry:getRegistry", TypeShape.of(GetRegistryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.harbor.HarborFunctions;
     * import com.pulumi.harbor.inputs.GetRegistryArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var main = HarborFunctions.getRegistry(GetRegistryArgs.builder()
     *             .name(&#34;test_docker_harbor&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;harborRegistryId&#34;, main.applyValue(getRegistryResult -&gt; getRegistryResult.id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetRegistryResult> getRegistryPlain(GetRegistryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("harbor:index/getRegistry:getRegistry", TypeShape.of(GetRegistryResult.class), args, Utilities.withVersion(options));
    }
}
