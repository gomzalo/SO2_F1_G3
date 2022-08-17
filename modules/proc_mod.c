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

static int writeFile(struct seq_file *archivo, void *v)
{

    seq_printf(archivo, "==============================\n");
    seq_printf(archivo, "=             OS2            =\n");
    seq_printf(archivo, "=            sopes2          =\n");
    seq_printf(archivo, "=           proc_mod         =\n");
    seq_printf(archivo, "==============================\n");

    return 0;
}

static int atOpen(struct inode *inode, struct file *file)
{
    return single_open(file, writeFile, NULL);
}

static const struct proc_ops ops = {
    .proc_open = atOpen,
    .proc_read = seq_read};

int proc_count(void)
{
    int i = 0;
    struct task_struct *thechild;
    
    for_each_process(thechild)
    {
        //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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
        if (thechild->pid==32)
        {
            /* code */
            //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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
        if (thechild->pid==1026)
        {
            /* code */
            //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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
        if (thechild->pid==0)
        {
            /* code */
            //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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
        if (thechild->pid==1)
        {
            /* code */
            //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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
        if (thechild->pid!=0 && thechild->pid!=1 && thechild->pid!=32 && thechild->pid!=1026)  
        {
            /* code */
            //pr_info("== %s [%d]\n", thechild->comm, thechild->pid);
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