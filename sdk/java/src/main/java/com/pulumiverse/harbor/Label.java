// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.harbor.LabelArgs;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.LabelState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * * Create a global label within harbor
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.Label;
 * import com.pulumi.harbor.LabelArgs;
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
 *         var main = new Label(&#34;main&#34;, LabelArgs.builder()        
 *             .color(&#34;#FF0000&#34;)
 *             .description(&#34;Description to for acceptance test&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * * Creates a label for project
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.Project;
 * import com.pulumi.harbor.Label;
 * import com.pulumi.harbor.LabelArgs;
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
 *         var mainProject = new Project(&#34;mainProject&#34;);
 * 
 *         var mainLabel = new Label(&#34;mainLabel&#34;, LabelArgs.builder()        
 *             .color(&#34;#FFFFFF&#34;)
 *             .description(&#34;Description for acceptance test&#34;)
 *             .projectId(mainProject.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Harbor label can be imported using the `label id` eg, `
 * 
 * ```sh
 *  $ pulumi import harbor:index/label:Label main /labels/1
 * ```
 * 
 *  `
 * 
 */
@ResourceType(type="harbor:index/label:Label")
public class Label extends com.pulumi.resources.CustomResource {
    @Export(name="color", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> color;

    public Output<Optional<String>> color() {
        return Codegen.optional(this.color);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output<String> scope;

    public Output<String> scope() {
        return this.scope;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Label(String name) {
        this(name, LabelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Label(String name, @Nullable LabelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Label(String name, @Nullable LabelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/label:Label", name, args == null ? LabelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Label(String name, Output<String> id, @Nullable LabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/label:Label", name, state, makeResourceOptions(options, id));
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
    public static Label get(String name, Output<String> id, @Nullable LabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Label(name, id, state, options);
    }
}
