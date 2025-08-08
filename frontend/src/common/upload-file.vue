<template>
    <div>
        <p class="text-2xl font-bold">Importar Datos del Excel</p>
        <p class="text-lg">
            Sube tu archivo Excel, mapea las columnas y valida los datos antes
            de importar
        </p>
    </div>
    <div class="card flex justify-center">
        <Stepper value="1" class="basis-[50rem]">
            <StepList>
                <Step value="1">Subir Archivo</Step>
                <Step value="2">Mapear Columnas</Step>
                <Step value="3">Validar Datos</Step>
            </StepList>
            <StepPanels>
                <StepPanel v-slot="{ activateCallback }" value="1">
                    <div class="flex flex-col h-48">
                        <div
                            class="border-2 border-surface-200 dark:border-surface-700 rounded bg-surface-50 dark:bg-surface-950 flex-auto flex justify-center items-center font-medium"
                        >
                            <FileUpload
                                ref="fileupload"
                                mode="basic"
                                name="file"
                                url="/api/upload"
                                accept=".xlsx"
                                v-model="file"
                                :maxFileSize="1000000"
                                @select="onSelect"
                                @error="onError"
                                @clear="onClear"
                            />
                        </div>
                    </div>
                    <div class="flex pt-6 justify-end">
                        <Button
                            label="Next"
                            icon="pi pi-arrow-right"
                            iconPos="right"
                            @click="validateFileColumns(activateCallback)"
                            :disabled="isNextDisabled"
                        />
                    </div>
                </StepPanel>
                <StepPanel v-slot="{ activateCallback }" value="2">
                    <div class="flex flex-col">
                        <div
                            class="mt-10 border-2 border-surface-200 dark:border-surface-700 rounded bg-surface-50 dark:bg-surface-950 flex flex-col p-4 gap-4"
                        >
                            <div>
                                <p class="font-black">Mapear Columnas</p>
                                <p class="">
                                    Relaciona las columnas de tu archivo con las
                                    columnas de la base de datos
                                </p>
                            </div>
                            <div>
                                <div class="grid grid-cols-3 gap-4">
                                    <div class="font-black">
                                        Columna del Archivo
                                    </div>
                                    <div class="font-black">
                                        Mapear a Columna de BD
                                    </div>
                                </div>
                                <Divider />
                                <div v-for="(item, index) in columnData">
                                    <div class="grid grid-cols-3 gap-4">
                                        <div class="">{{ item }}</div>
                                        <Select
                                            v-model="
                                                selectedMapping[index].database
                                            "
                                            :options="props.tableColumn"
                                            optionLabel="name"
                                            optionValue="value"
                                            placeholder="Selecciona una columna"
                                            class="w-full md:w-56"
                                        />
                                    </div>
                                    <Divider />
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="flex pt-6 justify-between">
                        <Button
                            label="Back"
                            severity="secondary"
                            icon="pi pi-arrow-left"
                            @click="activateCallback('1')"
                        />
                        <Button
                            label="Next"
                            icon="pi pi-arrow-right"
                            iconPos="right"
                            @click="updloadData(activateCallback)"
                        />
                    </div>
                </StepPanel>
                <StepPanel v-slot="{ activateCallback }" value="3">
                    <div class="flex flex-col h-48">
                        <div
                            class="mt-10 border-2 border-surface-200 dark:border-surface-700 rounded bg-surface-50 dark:bg-surface-950 flex flex-col p-4 gap-4"
                        >
                            <Message severity="success"
                                >¡Archivo procesado exitosamente!</Message
                            >
                            <Message severity="error"
                                >Se encontraron 2 errores que deben
                                corregirse</Message
                            >
                        </div>
                    </div>
                    <div class="flex pt-6 justify-between">
                        <Button
                            label="Back"
                            severity="secondary"
                            icon="pi pi-arrow-left"
                            @click="activateCallback('2')"
                        />
                        <!-- <Button
                            label="Next"
                            icon="pi pi-arrow-right"
                            iconPos="right"
                            @click="activateCallback('4')"
                        /> -->
                    </div>
                </StepPanel>
            </StepPanels>
        </Stepper>
    </div>
</template>
<script setup lang="ts">
import {
    Stepper,
    StepList,
    StepPanels,
    Step,
    StepPanel,
    FileUpload,
    Button,
    Divider,
    Select,
    Message,
} from "primevue";
import { ref, computed, watch } from "vue";
import { validateFile } from "../../api/validateDataApi";
import { uploadTeacherExcelData } from "../../api/profesoresApi";
import { uploadSubjectExcelData } from "../../api/cursosApi";
import { uploadRoomExcelData } from "../../api/salonesApi";
import type { FileUploadSelectEvent } from "primevue/fileupload";

const props = defineProps({
    items: {
        type: Array,
        default: [],
    },
    tableColumn: {
        type: Array,
        default: [],
    },
    type: {
        type: String,
        default: "",
    },
});

const selectedMapping = ref<Array<{ document: string; database: string }>>([]);
const stepperNumber = ref(1);
const columnData = ref([]);
const file = ref();

watch(
    () => columnData.value,
    (newItems: any) => {
        selectedMapping.value = newItems.map((item: any) => ({
            document: item,
            database: "",
        }));
    },
    { immediate: true },
);

watch(
    () => stepperNumber,
    async (newValue: any) => {
        if (newValue === 2) {
            await validateFile(file.value);
        }
    },
);

const fileupload = ref();
const fileName = ref("");
const hasUploadError = ref(false);
const isFileSelected = ref(false);

const onSelect = (event: FileUploadSelectEvent) => {
    file.value = event.files?.[0];
    isFileSelected.value = true;
};

const onError = () => {
    hasUploadError.value = true;
};

const onClear = () => {
    isFileSelected.value = false;
    hasUploadError.value = false;
};

const validateFileColumns = async (activateCallback: any) => {
    // fileupload.value.upload();
    console.log("UPLOAD");
    activateCallback("2");
    stepperNumber.value = 2;
    console.log("FILE", file.value);
    try {
        const result = await validateFile(file.value);
        columnData.value = result.columns;
        fileName.value = result.fileId;
        console.log("COLUMNDATA", result);
        hasUploadError.value = false;
    } catch (error: any) {
        console.log("ERROR", error);
        hasUploadError.value = true;
    }
};

const updloadData = async (activateCallback: any) => {
    activateCallback("3");
    if (props.type === "Teacher") {
        await uploadTeacherExcelData(fileName.value, selectedMapping.value);
    } else if (props.type === "Subject") {
        await uploadSubjectExcelData(fileName.value, selectedMapping.value);
    } else if (props.type === "Room") {
        await uploadRoomExcelData(fileName.value, selectedMapping.value);
    } else {
    }
};

const isNextDisabled = computed(
    () => hasUploadError.value || !isFileSelected.value,
);
</script>
