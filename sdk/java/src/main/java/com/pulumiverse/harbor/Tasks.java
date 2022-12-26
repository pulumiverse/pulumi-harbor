// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.harbor.TasksArgs;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.TasksState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.Tasks;
 * import com.pulumi.harbor.TasksArgs;
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
 *         var main = new Tasks(&#34;main&#34;, TasksArgs.builder()        
 *             .vulnerabilityScanPolicy(&#34;daily&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="harbor:index/tasks:Tasks")
public class Tasks extends com.pulumi.resources.CustomResource {
    @Export(name="vulnerabilityScanPolicy", refs={String.class}, tree="[0]")
    private Output<String> vulnerabilityScanPolicy;

    public Output<String> vulnerabilityScanPolicy() {
        return this.vulnerabilityScanPolicy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tasks(String name) {
        this(name, TasksArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tasks(String name, TasksArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tasks(String name, TasksArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/tasks:Tasks", name, args == null ? TasksArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Tasks(String name, Output<String> id, @Nullable TasksState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/tasks:Tasks", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Tasks get(String name, Output<String> id, @Nullable TasksState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Tasks(name, id, state, options);
    }
}
