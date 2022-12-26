// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.harbor.ProjectWebhookArgs;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.ProjectWebhookState;
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
 * import com.pulumi.harbor.Project;
 * import com.pulumi.harbor.ProjectWebhook;
 * import com.pulumi.harbor.ProjectWebhookArgs;
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
 *         var mainProjectWebhook = new ProjectWebhook(&#34;mainProjectWebhook&#34;, ProjectWebhookArgs.builder()        
 *             .address(&#34;https://webhook.domain.com&#34;)
 *             .projectId(mainProject.id())
 *             .notifyType(&#34;http&#34;)
 *             .eventsTypes(            
 *                 &#34;DELETE_ARTIFACT&#34;,
 *                 &#34;PULL_ARTIFACT&#34;,
 *                 &#34;PUSH_ARTIFACT&#34;,
 *                 &#34;DELETE_CHART&#34;,
 *                 &#34;DOWNLOAD_CHART&#34;,
 *                 &#34;UPLOAD_CHART&#34;,
 *                 &#34;QUOTA_EXCEED&#34;,
 *                 &#34;QUOTA_WARNING&#34;,
 *                 &#34;REPLICATION&#34;,
 *                 &#34;SCANNING_FAILED&#34;,
 *                 &#34;SCANNING_COMPLETED&#34;,
 *                 &#34;TAG_RETENTION&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="harbor:index/projectWebhook:ProjectWebhook")
public class ProjectWebhook extends com.pulumi.resources.CustomResource {
    /**
     * The address of the webhook
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The address of the webhook
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * authentication header for you the webhook
     * 
     */
    @Export(name="authHeader", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authHeader;

    /**
     * @return authentication header for you the webhook
     * 
     */
    public Output<Optional<String>> authHeader() {
        return Codegen.optional(this.authHeader);
    }
    /**
     * _ (Optional, string) A description of the webhook
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return _ (Optional, string) A description of the webhook
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * , To enable / disable the webhook. Default `true`
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return , To enable / disable the webhook. Default `true`
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `DELETE_CHART`, `DOWNLOAD_CHART`, `UPLOAD_CHART`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
     * 
     */
    @Export(name="eventsTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventsTypes;

    /**
     * @return ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `DELETE_CHART`, `DOWNLOAD_CHART`, `UPLOAD_CHART`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
     * 
     */
    public Output<List<String>> eventsTypes() {
        return this.eventsTypes;
    }
    /**
     * The name of the webhook that will be created in harbor.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the webhook that will be created in harbor.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The notification type either `http` or `slack`
     * 
     */
    @Export(name="notifyType", refs={String.class}, tree="[0]")
    private Output<String> notifyType;

    /**
     * @return The notification type either `http` or `slack`
     * 
     */
    public Output<String> notifyType() {
        return this.notifyType;
    }
    /**
     * The project id of the harbor that webhook related to.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project id of the harbor that webhook related to.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * checks the for validate SSL certificate.
     * 
     */
    @Export(name="skipCertVerify", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipCertVerify;

    /**
     * @return checks the for validate SSL certificate.
     * 
     */
    public Output<Optional<Boolean>> skipCertVerify() {
        return Codegen.optional(this.skipCertVerify);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectWebhook(String name) {
        this(name, ProjectWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectWebhook(String name, ProjectWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectWebhook(String name, ProjectWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/projectWebhook:ProjectWebhook", name, args == null ? ProjectWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectWebhook(String name, Output<String> id, @Nullable ProjectWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/projectWebhook:ProjectWebhook", name, state, makeResourceOptions(options, id));
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
    public static ProjectWebhook get(String name, Output<String> id, @Nullable ProjectWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectWebhook(name, id, state, options);
    }
}
