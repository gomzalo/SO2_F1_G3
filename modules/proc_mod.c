#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("sopes2");
MODULE_DESCRIPTION("Basic process information Linux module.");
MODULE_VERSION("0.01");

struct task_struct *task;
struct task_struct *task_child;
struct list_head *list;
int extra;
int extra2;

char * get_task_state(long state) {
    switch (state)
    {
    case 0:
        return "Running";
    case 1:
        return "Process";
    case 32:
        return "Zombie";
    case 1026:
        return "Stopped";
    default:
        return "Other";
    }
}

// static int writeFile(struct seq_file *archivo, void *v)
// {
//     seq_printf(archivo, "==============================\n");
//     seq_printf(archivo, "=             OS2            =\n");
//     seq_printf(archivo, "=            sopes2          =\n");
//     seq_printf(archivo, "=           proc_mod         =\n");
//     seq_printf(archivo, "==============================\n");
//     return 0;
// }

static int proc_llenar_archivo(struct seq_file *m, void *v) {
    
    #define K(x) ((x) << (PAGE_SHIFT - 10))

    seq_printf(m, "[\n");
    extra2 = 0;

    for_each_process(task) {

        if (extra2 == 0)
        {

            extra2 = 1;
        }
        else
        {

            seq_printf(m, ",");
        }
        
        seq_printf(m, "Process: %s\t PID:[%d]\t State: %ld\n", task->comm, task->pid, task->state);
        if (task->mm)
        {
            seq_printf(m, "\"mm\"  : %8lu, ", K(task->mm->total_vm)/2014);
        }
        else
        {
            seq_printf(m, "\"mm\"  : 0, ");
        }

        seq_printf(m, "\"sub\": [");

        extra = 0;

        list_for_each(list, &task->children)
        {

            if (extra == 0)
            {

                extra = 1;
            }
            else
            {

                seq_printf(m, ",");
            }

            task_child = list_entry(list, struct task_struct, sibling);
            seq_printf(m, "\n     { \"PID\" : %d, \"Nombre\" : \"%s\" , \"Estado\" : %ld , \"uid\" : %i,  ", task_child->pid, task_child->comm, task_child->state, task_child->cred->uid.val);
            if (task->mm)
            {
                seq_printf(m, "\"mm\"  : %8lu }", K(task->mm->total_vm)/1024);
            }
            else
            {
                seq_printf(m, "\"mm\"  : 0 }");
            }
        }

        extra = 0;
        seq_printf(m, "]\n}\n");
    }

    seq_printf(m, "\n]\n");
    return 0;
}

static int atOpen(struct inode *inode, struct file *file)
{
    return single_open(file, proc_llenar_archivo, NULL);
}

static const struct proc_ops ops = {
    .proc_open = atOpen,
    .proc_read = seq_read
};

int proc_count(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
        i++;
    }
    return i;
}

int proc_count_zombie(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        if (thechild->state==32)
        {
            /* code */
            // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
            i++;
        }
        
        
    }
    return i;
}

int proc_count_interrumpidos(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        if (thechild->state==1026)
        {
            /* code */
            // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
            i++;
        }
        
        
    }
    return i;
}

int proc_count_ejecucion(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        if (thechild->state==0)
        {
            /* code */
            // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
            i++;
        }
        
        
    }
    return i;
}

int proc_count_suspendidos(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        if (thechild->state==1)
        {
            /* code */
            // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
            i++;
        }
        
        
    }
    return i;
}

int proc_count_detenidos(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        if (thechild->state!=0 && thechild->state!=1 && thechild->state!=32 && thechild->state!=1026)  
        {
            /* code */
            // pr_info("== %s [%d]\n", thechild->comm, thechild->state);
            i++;
        }
        
        
    }
    return i;
}


static int load_module(void)
{
    printk(KERN_INFO "Total processes: %d .\n", proc_count());
    printk(KERN_INFO "Total running processes: %d .\n", proc_count_ejecucion());
    printk(KERN_INFO "Total zombie processes: %d .\n", proc_count_zombie());
    printk(KERN_INFO "Total interrumpidos processes: %d .\n", proc_count_interrumpidos());
     printk(KERN_INFO "Total suspendidos processes: %d .\n", proc_count_suspendidos());
    printk(KERN_INFO "Total detenidos processes: %d .\n", proc_count_detenidos());

    proc_create("proc_mod", 0, NULL, &ops);
    return 0;
}

static void unload_module(void)
{
    printk(KERN_INFO "Goodbye!\n");

    remove_proc_entry("proc_mod", NULL);
}

module_init(load_module);
module_exit(unload_module);