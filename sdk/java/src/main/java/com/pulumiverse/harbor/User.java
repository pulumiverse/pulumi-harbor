// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.harbor.UserArgs;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.UserState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.User;
 * import com.pulumi.harbor.UserArgs;
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
 *         var main = new User(&#34;main&#34;, UserArgs.builder()        
 *             .email(&#34;john@smith.com&#34;)
 *             .fullName(&#34;John Smith&#34;)
 *             .password(&#34;Password12345!&#34;)
 *             .username(&#34;john&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * An internal user harbor user can be imported using the `user id` eg, `
 * 
 * ```sh
 *  $ pulumi import harbor:index/user:User main /users/19
 * ```
 * 
 *  `
 * 
 */
@ResourceType(type="harbor:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    @Export(name="admin", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> admin;

    public Output<Optional<Boolean>> admin() {
        return Codegen.optional(this.admin);
    }
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    public Output<String> email() {
        return this.email;
    }
    @Export(name="fullName", refs={String.class}, tree="[0]")
    private Output<String> fullName;

    public Output<String> fullName() {
        return this.fullName;
    }
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    public Output<String> password() {
        return this.password;
    }
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
